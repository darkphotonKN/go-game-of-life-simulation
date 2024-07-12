package game

import (
	"fmt"
)

type board struct {
	Width  int
	Height int
	Grid   [][]bool // presents alive or dead cell
}

// new board
func NewBoard() *board {
	b := &board{}

	b.initGrid(10)

	return b
}

// initialize grid
func (b *board) initGrid(width int) {
	// create multi-dimensional array to represent grid
	b.Grid = make([][]bool, width)

	// fill it with falses to start with
	for rowIndex, _ := range b.Grid {
		// create inner slice size

		b.Grid[rowIndex] = make([]bool, width)
		for colIndex, _ := range b.Grid[rowIndex] {
			b.Grid[rowIndex][colIndex] = false
		}
	}

	b.Width = width
	b.Height = width
}

// visualize grid
func (b *board) Visualize() {

	// extra row at the top for board top visualization
	for i := 0; i < b.Width; i++ {
		fmt.Print("  - ")
	}

	fmt.Println()

	// iterate through mutidimensional array
	for _, row := range b.Grid {
		fmt.Print("|")
		for _, cell := range row {

			// if cell is empty render space
			if !cell {
				fmt.Print("  ")
			} else {
				// else render an x to show an alive cell
				fmt.Print(" x")
			}
			// fmt.Println("  -")
			fmt.Print(" |")
		}
		fmt.Println("  ")
	}

	// extra row at the bottom for board bottom visualization
	for i := 0; i < b.Width; i++ {
		fmt.Print("  - ")
	}

}
