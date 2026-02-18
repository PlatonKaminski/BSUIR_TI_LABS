import { useState } from 'react';
import './App.css';
import { EncryptV, DecryptV, EncryptM, DecryptM } from "../wailsjs/go/main/App";

function App() {
    const [sourceText, setSourceText] = useState('');
    const [resultText, setResultText] = useState('');
    const [keyText, setKeyText] = useState('');
    const [selectedMethod, setSelectedMethod] = useState('vigenere');
    const GRILLE_SIZE = 4;

    const handleMethodChange = (e) => {
        setSelectedMethod(e.target.value);
        setResultText('');
        setKeyText('');
        setSourceText('');
    };

    const handleSourceChange = (e) => {
        const value = e.target.value;
        let filteredValue = value;

        if (selectedMethod === 'vigenere') {
            filteredValue = value.replace(/[^а-яА-ЯёЁ\s]/g, '');
        } else if (selectedMethod === 'grille') {
            filteredValue = value.replace(/[^a-zA-Z\s]/g, '');

        }

        setSourceText(filteredValue);
        setResultText('');
    };

    const handleKeyChange = (e) => {
        const value = e.target.value;
        let filteredValue = value;

        if (selectedMethod === 'vigenere') {
            filteredValue = value.replace(/[^а-яА-ЯёЁ]/g, '');
        }

        setKeyText(filteredValue);
        setResultText('');
    };

    const handleEncrypt = () => {
        if (selectedMethod === 'vigenere') {
            if (!sourceText.trim() || !keyText.trim()) {
                setResultText('⚠️ Введите текст и ключ на русском!');
                return;
            }

            EncryptV(sourceText, keyText).then(setResultText).catch(error => {
                setResultText('⚠️ Ошибка шифрования');
                console.error(error);
            });
        } else if (selectedMethod === 'grille') {
            if (!sourceText.trim()) {
                setResultText('⚠️ Введите текст на английском!');
                return;
            }



            EncryptM(sourceText, GRILLE_SIZE.toString()).then(setResultText).catch(error => {
                setResultText('⚠️ Ошибка шифрования');
                console.error(error);
            });
        }
    };

    const handleDecrypt = () => {
        if (selectedMethod === 'vigenere') {
            if (!sourceText.trim() || !keyText.trim()) {
                setResultText('⚠️ Введите текст и ключ на русском!');
                return;
            }
            DecryptV(sourceText, keyText).then(setResultText).catch(error => {
                setResultText('⚠️ Ошибка дешифрования');
                console.error(error);
            });
        } else if (selectedMethod === 'grille') {
            if (!sourceText.trim()) {
                setResultText('⚠️ Введите текст на английском!');
                return;
            }

            const requiredLength = GRILLE_SIZE * GRILLE_SIZE;
            if (sourceText.length !== requiredLength) {
                setResultText(`⚠️ Зашифрованный текст должен быть длиной ${requiredLength} символов!`);
                return;
            }

            DecryptM(sourceText, GRILLE_SIZE.toString()).then(setResultText).catch(error => {
                setResultText('⚠️ Ошибка дешифрования');
                console.error(error);
            });
        }
    };

    const handleClear = () => {
        setSourceText('');
        setResultText('');
        setKeyText('');
    };

    const handleReadFile = () => {
        const input = document.createElement('input');
        input.type = 'file';
        input.accept = '.txt,.text,.md,.csv';

        input.onchange = (e) => {
            const file = e.target.files[0];
            if (!file) return;

            const reader = new FileReader();
            reader.onload = (event) => {
                const content = event.target.result;
                let filteredContent = content;

                if (selectedMethod === 'vigenere') {
                    filteredContent = content.replace(/[^а-яА-ЯёЁ\s]/g, '');
                } else if (selectedMethod === 'grille') {
                    filteredContent = content.replace(/[^a-zA-Z\s]/g, '');
                    if (filteredContent.length > GRILLE_SIZE * GRILLE_SIZE) {
                        filteredContent = filteredContent.slice(0, GRILLE_SIZE * GRILLE_SIZE);
                    }
                }

                setSourceText(filteredContent);
                setResultText('');
            };
            reader.readAsText(file);
        };

        input.click();
    };

    return (
        <div id="App">
            <div className="app-header">
                <div className="app-title">МЕТОДЫ ШИФРОВАНИЯ</div>
            </div>

            {/* Выбор метода шифрования */}
            <div className="input-group">
                <div className="label">Выберите метод шифрования</div>
                <div className="radio-group">
                    <label className="radio-label">
                        <input
                            type="radio"
                            name="cipherMethod"
                            value="vigenere"
                            checked={selectedMethod === 'vigenere'}
                            onChange={handleMethodChange}
                        />
                        <span style={{color: 'white'}}>Шифр Виженера</span>
                    </label>
                    <label className="radio-label">
                        <input
                            type="radio"
                            name="cipherMethod"
                            value="grille"
                            checked={selectedMethod === 'grille'}
                            onChange={handleMethodChange}
                        />
                        <span style={{color: 'white'}}>Шифр поворотной решетки (4x4)</span>
                    </label>
                </div>
            </div>

            {/* Поле для ключа */}
            <div className="input-group">
                <div className="label">Ключ шифрования</div>
                <input
                    className={`input ${selectedMethod === 'grille' ? 'input-disabled' : ''}`}
                    value={keyText}
                    onChange={handleKeyChange}
                    disabled={selectedMethod === 'grille'}
                    autoComplete="off"
                    type="text"
                    placeholder={selectedMethod === 'vigenere'
                        ? "Введите ключ на русском..."
                        : "Ключ не требуется для этого метода"}
                />
                <div className="input-hint">
                    {selectedMethod === 'vigenere'
                        ? "Только русские буквы"
                        : "Поле недоступно для шифра поворотной решетки"}
                </div>
            </div>

            <div className="input-group">
                <div className="label">Исходный текст</div>
                <input
                    className="input"
                    value={sourceText}
                    onChange={handleSourceChange}
                    autoComplete="off"
                    type="text"
                    placeholder={selectedMethod === 'vigenere'
                        ? "Введите текст на русском..."
                        : `Введите текст на английском`}
                />
                <div className="input-hint">
                    {selectedMethod === 'vigenere'
                        ? "Только русские буквы и пробелы"
                        : `Только английские буквы`}
                </div>
                {selectedMethod === 'grille' && (
                    <div className="input-counter">
                    </div>
                )}
            </div>

            <div className="input-group">
                <div className="label">Результирующий текст</div>
                <input
                    className="input"
                    value={resultText}
                    readOnly
                    type="text"
                    placeholder="Результат..."
                />
            </div>

            <div className="button-group">
                <button className="btn btn-encrypt" onClick={handleEncrypt}>
                    Шифровать
                </button>
                <button className="btn btn-decrypt" onClick={handleDecrypt}>
                    Дешифровать
                </button>
                <button className="btn btn-file" onClick={handleReadFile}>
                    Прочитать из файла
                </button>
                <button className="btn btn-clear" onClick={handleClear}>
                    Очистить
                </button>
            </div>
        </div>
    );
}

export default App;