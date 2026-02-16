package main

import (
	"context"
	"strings"
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

const SIZE_ALPHABET = 33

func (a *App) Encrypt(str string, key string) string {
	i := 0
	resultStr := ""

	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
	key = strings.ToUpper(strings.ReplaceAll(key, " ", ""))

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

func (a *App) Decrypt(str string, key string) string {
	i := 0
	resultStr := ""

	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
	key = strings.ToUpper(strings.ReplaceAll(key, " ", ""))

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
