package main

import "fmt"

const (
	Tableline = "======================="
	rowFormat = "|%c|\n"
	inline    = "__"
)

func display(board [10][10]rune) {
	fmt.Println(Tableline)
	for _, row := range board {
		for _, column := range row {
			if row == 0 {
				fmt.Print("|%c", column)
			} else {
				fmt.Printf(rowFormat, " ")
			}
		}
		fmt.Println(Tableline)

	}
}

func main() {
	var board [10][10]rune
	//黑方棋子
	board[1][1] = 'r'
	board[1][2] = 'n'
	board[1][3] = 'b'
	board[1][4] = 'q'
	board[1][5] = 'k'
	board[1][6] = 'b'
	board[1][7] = 'n'
	board[1][8] = 'r'
	//白方棋子
	board[9][1] = 'R'
	board[9][2] = 'N'
	board[9][3] = 'B'
	board[9][4] = 'Q'
	board[9][5] = 'K'
	board[9][6] = 'B'
	board[9][7] = 'N'
	board[9][8] = 'R'

	/*	for column := range board[2] {
		board[2][column] = 'p'
		board[6][column] = 'p'
	}*/

	display(board)
}
