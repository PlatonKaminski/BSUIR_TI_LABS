package main

import (
	"context"
	"math/rand"
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

func (a *App) Encrypt(str string) string {
	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
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

func (a *App) Decrypt(str string) string {
	str = strings.ToUpper(strings.ReplaceAll(str, " ", ""))
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
