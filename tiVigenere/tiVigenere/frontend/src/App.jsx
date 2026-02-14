import { useState } from 'react';
import './App.css';
import { Encrypt, Decrypt } from "../wailsjs/go/main/App";

function App() {
    const [sourceText, setSourceText] = useState('');
    const [resultText, setResultText] = useState('');
    const [keyText, setKeyText] = useState(''); // состояние для ключа

    const handleSourceChange = (e) => {
        const value = e.target.value;
        // Разрешены только a-z, A-Z и пробел
        const onlyEnglish = value.replace(/[^a-zA-Z\s]/g, '');
        setSourceText(onlyEnglish);
    };

    const handleKeyChange = (e) => {
        const value = e.target.value;
        // Для ключа тоже только английские буквы
        const onlyEnglish = value.replace(/[^a-zA-Z]/g, '');
        setKeyText(onlyEnglish);
    };

    const handleEncrypt = () => {
        if (!sourceText.trim() || !keyText.trim()) {
            setResultText('⚠️ Введите текст и/или ключ!');
            return;
        }
        // Передаем и текст, и ключ
        Encrypt(sourceText, keyText).then(setResultText);
    };
    
    const handleDecrypt = () => {
        if (!sourceText.trim() || !keyText.trim()) {
            setResultText('⚠️ Введите текст и/или ключ!');
            return;
        }
        Decrypt(sourceText, keyText).then(setResultText);
    };

    const handleClear = () => {
        setSourceText('');
        setResultText('');
        setKeyText('');
    };

    // Функция для чтения из файла
    const handleReadFile = () => {
        // Создаем скрытый input элемент
        const input = document.createElement('input');
        input.type = 'file';
        input.accept = '.txt,.text,.md,.csv'; // разрешенные расширения
        
        input.onchange = (e) => {
            const file = e.target.files[0];
            if (!file) return;
            
            const reader = new FileReader();
            reader.onload = (event) => {
                const content = event.target.result;
                // Фильтруем только английские буквы и пробелы
                const onlyEnglish = content.replace(/[^a-zA-Z\s]/g, '');
                setSourceText(onlyEnglish);
                
            }
            reader.readAsText(file);
        };
        
        input.click();
    };

    return (
        <div id="App">
            <div className="app-header">
                <div className="app-title">МЕТОД ВИЖЕНЕРА</div>
                <div className="app-subtitle">с прогрессивным ключом</div>
            </div>
            {/* НОВОЕ: Поле для ключа */}
            <div className="input-group">
                <div className="label">Ключ шифрования</div>
                <input 
                    className="input" 
                    value={keyText}
                    onChange={handleKeyChange} 
                    autoComplete="off" 
                    type="text"
                    placeholder="Введите ключ..."
                />
                <div className="input-hint">Только английские буквы</div>
            </div>

            <div className="input-group">
                <div className="label">Исходный текст</div>
                <input 
                    className="input" 
                    value={sourceText}
                    onChange={handleSourceChange} 
                    autoComplete="off" 
                    type="text"
                    placeholder="Введите текст..."
                />
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
                {/* НОВОЕ: Кнопка чтения из файла */}
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