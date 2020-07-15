package main

import (
	"os"

	"github.com/huiyeon5/tic-tac-go/pkg/board"
)

func main() {
	tttBoard := board.NewBoard()
	tttBoard.Display(os.Stdin)
	tttBoard.Mark("X", 0)
	tttBoard.Mark("X", 2)
	tttBoard.Mark("X", 4)
	tttBoard.Display(os.Stdin)
}
