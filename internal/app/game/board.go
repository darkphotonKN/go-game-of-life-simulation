package game

import (
	"fmt"
	"math/rand"
	"time"
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

	b.seedGrid(5)

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

// seed grid

func (b *board) seedGrid(x int) {
	// randomly pick x number of spots

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < x; i++ {

		randomRow := random.Intn(b.Width)  // Generate random numbers between 0 and x
		randomCol := random.Intn(b.Height) // Generate random numbers between 0 and x

		b.Grid[randomRow][randomCol] = true
	}
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
