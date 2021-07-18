package main

import "fmt"

const (
	Tableline = "=================="
	rowFormat = "%c|\n"
	inline    = "_"
)

func display(board [8][8]string) {
	fmt.Println(Tableline)
	for _, row := range board {
		fmt.Printf("|")
		for _, column := range row {
			fmt.Printf("%v", column)
			if column == "" {
				fmt.Printf(" ")
			}
			fmt.Printf("|")
		}
		fmt.Printf("\n")
		fmt.Printf("------------------\n")
	}

	fmt.Println(Tableline)

}

func main() {
	var board [8][8]string
	//黑方棋子
	board[0][0] = "r"
	board[0][1] = "n"
	board[0][2] = "b"
	board[0][3] = "q"
	board[0][4] = "k"
	board[0][5] = "b"
	board[0][6] = "n"
	board[0][7] = "r"
	//白方子
	board[7][0] = "R"
	board[7][1] = "N"
	board[7][2] = "B"
	board[7][3] = "Q"
	board[7][4] = "K"
	board[7][5] = "B"
	board[7][6] = "N"
	board[7][7] = "R"
	display(board)
}
