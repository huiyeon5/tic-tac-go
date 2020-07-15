package board

import (
	"errors"
	"fmt"
	"io"
)

// Board - Tic Tac Toe Board
type Board struct {
	values [9]string
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

// ErrOutOfBounds - Error for accessing position out of bounds
var ErrOutOfBounds = errors.New("Position out of bounds")

// ErrMarkedPos - Error for marking already marked position
var ErrMarkedPos = errors.New("Position has been marked")
