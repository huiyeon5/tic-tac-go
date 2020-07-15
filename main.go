package main

import (
	"os"

	util "github.com/huiyeon5/tic-tac-go/pkg"
)

func main() {
	tttBoard := util.NewBoard()
	tttBoard.Display(os.Stdin)
	tttBoard.Mark("X", 0)
	tttBoard.Mark("X", 2)
	tttBoard.Mark("X", 4)
	tttBoard.Display(os.Stdin)
}
