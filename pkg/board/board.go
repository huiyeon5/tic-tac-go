package board

import (
	"errors"
	"fmt"
	"io"
)

const (
	// Win - Player won
	Win = iota
	// Continue - No one won/lost
	Continue
	// Draw - It is a draw
	Draw
)

// Board - Tic Tac Toe Board
type Board struct {
	values [9]string
	writer io.Writer
}

// Init - Initializes the board with empty spaces
func (b *Board) Init() {
	for i := 0; i < 9; i++ {
		b.values[i] = " "
	}
}

// NewBoard - Creates and Returns an initialized board
func NewBoard() Board {
	board := Board{}
	board.Init()
	return board
}

// Display - Displays the Board to the stream
func (b *Board) Display(writer io.Writer) {
	for i := 0; i < 9; i++ {
		fmt.Fprint(writer, b.values[i])
		if i == 8 {
			fmt.Fprintln(writer)
			continue
		}

		if i == 2 || i == 5 {
			fmt.Fprintln(writer)
			continue
		}

		fmt.Fprint(writer, "|")
	}
}

// Mark - Marks a position on board
func (b *Board) Mark(val string, pos int) error {
	if pos < 0 || pos > 8 {
		return ErrOutOfBounds
	}

	if b.values[pos] != " " {
		return ErrMarkedPos
	}

	b.values[pos] = val
	return nil
}

// VerifyBoard - Verifies board for action
func (b *Board) VerifyBoard(player string) int {
	if checkRow(b, player) {
		return Win
	}

	if checkCol(b, player) {
		return Win
	}

	if checkDiagonal(b, player) {
		return Win
	}

	if checkFull(b) {
		return Draw
	}

	return Continue
}

func checkRow(b *Board, player string) bool {
	for i := 2; i < 9; i += 3 {
		if b.values[i] == b.values[i-1] && b.values[i-1] == b.values[i-2] && b.values[i-2] == player {
			return true
		}
	}
	return false
}

func checkCol(b *Board, player string) bool {
	for i := 0; i < 3; i++ {
		if b.values[i] == b.values[i+3] && b.values[i+3] == b.values[i+6] && b.values[i+6] == player {
			return true
		}
	}
	return false
}

func checkDiagonal(b *Board, player string) bool {
	if b.values[0] == b.values[4] && b.values[4] == b.values[8] && b.values[8] == player {
		return true
	}

	if b.values[2] == b.values[4] && b.values[4] == b.values[6] && b.values[6] == player {
		return true
	}
	return false
}

func checkFull(b *Board) bool {
	for _, val := range b.values {
		if val == " " {
			return false
		}
	}
	return true
}

// ErrOutOfBounds - Error for accessing position out of bounds
var ErrOutOfBounds = errors.New("Position out of bounds")

// ErrMarkedPos - Error for marking already marked position
var ErrMarkedPos = errors.New("Position has been marked")
