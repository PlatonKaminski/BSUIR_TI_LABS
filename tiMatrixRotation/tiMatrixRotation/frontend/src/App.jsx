import { useState } from 'react';
import './App.css';
import { Encrypt, Decrypt } from "../wailsjs/go/main/App";

function App() {
    const [sourceText, setSourceText] = useState('');
    const [resultText, setResultText] = useState('');

    const handleSourceChange = (e) => {
    const value = e.target.value;
    // Разрешены только a-z, A-Z и пробел
    const onlyEnglish = value.replace(/[^a-zA-Z\s]/g, '');
    setSourceText(onlyEnglish);
};

    const handleEncrypt = () => {
        if (!sourceText.trim()) {
            return;
        }
        Encrypt(sourceText).then(setResultText);
    };
    
    const handleDecrypt = () => {
        if (!sourceText.trim()) {
            return;
        }
        const lettersOnly = sourceText.replace(/\s+/g, '');
        if (lettersOnly.length < 16) {
            setResultText('Недостаточно символов. Нужно  16 букв.');
            return;
        }
        Decrypt(sourceText).then(setResultText);
    };

    const handleClear = () => {
        setSourceText('');
        setResultText('');
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
                <div className="app-title">МЕТОД ПОВОРОТА РЕШЕТКИ</div>
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