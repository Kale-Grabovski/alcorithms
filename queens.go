package main

import (
	"fmt"
	"math/rand"
	"time"
)

const queen = "Q"

func boardValid(board [][]string) bool {
	var q, i2, k2 int

	// rows
	for i := 0; i < 8; i++ {
		q = 0
		for k := 0; k < 8; k++ {
			if board[i][k] != "" {
				q++
			}
		}
		if q > 1 {
			return false
		}
	}

	// cols
	for i := 0; i < 8; i++ {
		q = 0
		for k := 0; k < 8; k++ {
			if board[k][i] != "" {
				q++
			}
		}
		if q > 1 {
			return false
		}
	}

	// diags
	for i := 0; i < 8; i++ {
		for k := 0; k < 8; k++ {
			if board[i][k] == "" {
				continue
			}

			// top-left
			i2, k2 = i, k
			for {
				i2--
				k2--
				if i2 < 0 || k2 < 0 {
					break
				}
				if board[i2][k2] != "" {
					return false
				}
			}

			// top-right
			i2, k2 = i, k
			for {
				i2--
				k2++
				if i2 < 0 || k2 == 8 {
					break
				}
				if board[i2][k2] != "" {
					return false
				}
			}

			// bottom-left
			i2, k2 = i, k
			for {
				i2++
				k2--
				if i2 == 8 || k2 < 0 {
					break
				}
				if board[i2][k2] != "" {
					return false
				}
			}

			// bottom-right
			i2, k2 = i, k
			for {
				i2++
				k2++
				if i2 == 8 || k2 == 8 {
					break
				}
				if board[i2][k2] != "" {
					return false
				}
			}
		}
	}

	return true
}

func printBoard(board [][]string) {
	for _, v := range board {
		for _, v2 := range v {
			if v2 == "" {
				fmt.Print(". ")
			} else {
				fmt.Print(queen + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func placeQueens() [][]string {
	rand.Seed(time.Now().Unix())

	board := make([][]string, 8, 8)
	for i := 0; i < 8; i++ {
		board[i] = make([]string, 8, 8)
	}

	row := 0

OUTER:
	for {
		for i := 0; i < 8; i++ {
			p := rand.Intn(8)
			board[row][p] = queen
			if boardValid(board) {
				if row == 7 {
					break OUTER
				}
				row++
				continue OUTER
			}
			board[row][p] = ""
		}

		// backtrack
		for i := 0; i < 8; i++ {
			board[row][i] = ""
			board[row-1][i] = ""
		}
		row--
	}

	return board
}

func main() {
	board := placeQueens()
	printBoard(board)
}
