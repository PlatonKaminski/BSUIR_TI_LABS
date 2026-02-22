package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

const SIZE = 4
const SIZE_STRING = 16
const SIZE_ALPHABET = 33

func filterRussianText(text string) string {
	var result strings.Builder
	for _, r := range text {
		if unicode.Is(unicode.Cyrillic, r) || r == ' ' {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func filterEnglishText(text string) string {
	var result strings.Builder
	for _, r := range text {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || r == ' ' {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func filterRussianKey(text string) string {
	var result strings.Builder
	for _, r := range text {
		if unicode.Is(unicode.Cyrillic, r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func (a *App) EncryptV(str string, key string) string {
	str = filterRussianText(str)
	key = filterRussianKey(key)

	i := 0
	resultStr := ""

	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
	key = strings.ToUpper(strings.ReplaceAll(key, " ", ""))

	if len(str) == 0 {
		return "Ошибка: нет допустимых русских букв для шифрования"
	}
	if len(key) == 0 {
		return "Ошибка: нет допустимых русских букв в ключе"
	}

	strRunes := []rune(str)
	keyRunes := []rune(key)

	for len(strRunes) > len(keyRunes) {
		keyRunes = append(keyRunes, keyRunes[i]+1)
		i++
	}

	for i := 0; i < len(strRunes); i++ {
		strIdx := getRuneIndex(strRunes[i])
		keyIdx := getRuneIndex(keyRunes[i])
		encryptedIdx := (strIdx + keyIdx) % SIZE_ALPHABET
		resultStr += string(getRuneByIndex(encryptedIdx))
	}

	return resultStr
}

func (a *App) DecryptV(str string, key string) string {
	str = filterRussianText(str)
	key = filterRussianKey(key)

	i := 0
	resultStr := ""

	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
	key = strings.ToUpper(strings.ReplaceAll(key, " ", ""))

	if len(str) == 0 {
		return "Ошибка: нет допустимых русских букв для дешифрования"
	}
	if len(key) == 0 {
		return "Ошибка: нет допустимых русских букв в ключе"
	}

	strRunes := []rune(str)
	keyRunes := []rune(key)

	for len(strRunes) > len(keyRunes) {
		keyRunes = append(keyRunes, keyRunes[i]+1)
		i++
	}

	for i := 0; i < len(strRunes); i++ {
		strIdx := getRuneIndex(strRunes[i])
		keyIdx := getRuneIndex(keyRunes[i])

		decryptedIdx := (strIdx - keyIdx + SIZE_ALPHABET) % SIZE_ALPHABET
		resultStr += string(getRuneByIndex(decryptedIdx))
	}

	return resultStr
}

func getRuneIndex(r rune) int {
	alphabet := []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")

	for i, letter := range alphabet {
		if letter == r {
			return i
		}
	}
	return 0
}

func getRuneByIndex(idx int) rune {
	alphabet := []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")

	if idx >= 0 && idx < len(alphabet) {
		return alphabet[idx]
	}
	return 'А'
}

func speenMatrix(matrix *[SIZE][SIZE]byte) [SIZE][SIZE]byte {
	var temparr [SIZE][SIZE]byte = [SIZE][SIZE]byte{}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			temparr[i][j] = matrix[j][SIZE-i-1]
		}
	}
	*matrix = temparr
	return *matrix
}

func (a *App) EncryptM(str string) string {
	str = filterEnglishText(str)
	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
	if len(str) == 0 {
		return "Ошибка: нет допустимых английских букв для шифрования"
	}

	var resultStr string = ""
	for len(str) < SIZE_STRING {
		c := byte(rand.Intn(26) + 65)
		str += string(c)
	}
	var matrix [SIZE][SIZE]byte = [SIZE][SIZE]byte{}
	i := 0
	positions := [][2]int{{0, 0}, {1, 3}, {2, 2}, {3, 1}}

	for j := 0; j < SIZE; j++ {
		for _, pos := range positions {
			matrix[pos[0]][pos[1]] = str[i]
			i++
		}
		speenMatrix(&matrix)
	}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			resultStr += string(matrix[i][j])
		}
	}
	return resultStr
}

func (a *App) DecryptM(str string) string {
	str = filterEnglishText(str)
	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))

	if len(str) < SIZE_STRING {
		return "Ошибка: недостаточно английских букв для дешифрования"
	}
	if len(str) > SIZE_STRING {
		str = str[:SIZE_STRING]
	}

	var resultStr string = ""
	k := 0
	var matrix [SIZE][SIZE]byte = [SIZE][SIZE]byte{}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			matrix[i][j] = str[k]
			k++
		}
	}

	positions := [][2]int{{0, 0}, {1, 3}, {2, 2}, {3, 1}}

	for j := 0; j < SIZE; j++ {
		for _, pos := range positions {
			resultStr += string(matrix[pos[0]][pos[1]])
		}
		speenMatrix(&matrix)
	}

	return resultStr
}

func (a *App) SaveFile() (string, error) {
	filePath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Сохранить файл",
		DefaultFilename: "result.txt",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Текстовые файлы",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Все файлы",
				Pattern:     "*.*",
			},
		},
	})
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (a *App) WriteFile(filePath string, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
