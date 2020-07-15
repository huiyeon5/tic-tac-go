package board

import (
	"bytes"
	"testing"
)

func TestDisplay(t *testing.T) {
	t.Run("Display Empty Board", func(t *testing.T) {
		board := NewBoard()
		buffer := &bytes.Buffer{}
		board.Display(buffer)

		got := buffer.String()
		want := ` | | 
 | | 
 | | 
`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Display Marked Board", func(t *testing.T) {
		board := NewBoard()
		board.Mark("X", 4)
		board.Mark("O", 1)
		buffer := &bytes.Buffer{}
		board.Display(buffer)

		got := buffer.String()
		want := ` |O| 
 |X| 
 | | 
`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	for _, v := range board.values {
		if v != " " {
			t.Errorf("want %s but got %s", " ", v)
		}
	}
}

func TestMark(t *testing.T) {
	t.Run("Mark Empty Position", func(t *testing.T) {
		board := NewBoard()
		err := board.Mark("X", 5)

		assertNoError(t, err)

		for i, v := range board.values {
			if i == 5 && v != "X" {
				t.Errorf("Want %s at pos %d but got %s", "X", 5, v)
			}
		}
	})

	t.Run("Mark out of bounds position", func(t *testing.T) {
		board := NewBoard()
		err := board.Mark("X", 10)

		assertError(t, err, ErrOutOfBounds)
	})

	t.Run("Mark at already marked position", func(t *testing.T) {
		board := NewBoard()
		err := board.Mark("X", 6)
		err = board.Mark("X", 6)

		assertError(t, err, ErrMarkedPos)
	})
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
