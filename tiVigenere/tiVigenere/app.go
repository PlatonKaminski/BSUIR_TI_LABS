package main

import (
	"context"
	"strings"
)

const SIZE_ALPHABET = 26

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

func (a *App) Encrypt(str string, key string) string {
	i := 0
	resultStr := ""
	for len(str) > len(key) {
		key += string(key[i] + 1)
		i++
	}
	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
	key = strings.ToUpper(strings.ReplaceAll(key, " ", ""))
	for i := 0; i < len(str); i++ {
		resultStr += string(((str[i]-'A')+(key[i]-'A'))%SIZE_ALPHABET + 'A')
	}

	return resultStr
}

func (a *App) Decrypt(str string, key string) string {
	i := 0
	resultStr := ""
	for len(str) > len(key) {
		key += string(key[i] + 1)
		i++
	}
	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
	key = strings.ToUpper(strings.ReplaceAll(key, " ", ""))
	for i := 0; i < len(str); i++ {
		resultStr += string(((str[i]-'A')-(key[i]-'A')+SIZE_ALPHABET)%SIZE_ALPHABET + 'A')
	}

	return resultStr
}
