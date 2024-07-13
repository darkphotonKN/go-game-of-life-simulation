package game

import (
	"fmt"
	"math/rand"
	"time"
)

/* Rules
* Any live cell with fewer than two live neighbours dies, as if by underpopulation.
* Any live cell with two or three live neighbours lives on to the next generation.
* Any live cell with more than three live neighbours dies, as if by overpopulation.
* Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
 */

type board struct {
	Width  int
	Height int
	Cell   [][]bool // represents alive or dead cell
}

// new board
func NewBoard() *board {
	b := &board{}

	b.initCell(10)

	b.seedCell(30)

	return b
}

// initialize grid
func (b *board) initCell(width int) {
	// create multi-dimensional array to represent grid
	b.Cell = make([][]bool, width)

	// fill it with falses to start with
	for rowIndex, _ := range b.Cell {
		// create inner slice size

		b.Cell[rowIndex] = make([]bool, width)
		for colIndex, _ := range b.Cell[rowIndex] {
			b.Cell[rowIndex][colIndex] = false
		}
	}

	b.Width = width
	b.Height = width
}

// seed grid
func (b *board) seedCell(x int) {
	// TODO: remove after testing
	// b.Cell[9][9] = true
	// b.Cell[8][9] = true
	// b.Cell[8][8] = true
	// b.Cell[0][1] = true
	// b.Cell[0][0] = true
	// b.Cell[1][1] = true

	// randomly pick x number of spots

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < x; i++ {

		randomRow := random.Intn(b.Width)  // Generate random numbers between 0 and x
		randomCol := random.Intn(b.Height) // Generate random numbers between 0 and x

		b.Cell[randomRow][randomCol] = true
	}
}

// visualize grid
func (b *board) visualize() {

	// extra row at the top for board top visualization
	for i := 0; i < b.Width; i++ {
		fmt.Print("  - ")
	}

	fmt.Println()

	// iterate through mutidimensional array
	for _, row := range b.Cell {
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

/* Rules
* Any live cell with fewer than two live neighbours dies, as if by underpopulation.
* Any live cell with two or three live neighbours lives on to the next generation.
* Any live cell with more than three live neighbours dies, as if by overpopulation.
* Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
 */

// handle simulation of the generations
func (b *board) Simulation() {
	generations := 1

	for {
		fmt.Println()
		fmt.Println("Generation:", generations)
		fmt.Println()
		generations++

		b.nextGeneration()

		b.visualize()

		time.Sleep(time.Millisecond * 900)

	}
}

// advance to next generation
func (b *board) nextGeneration() {

	for row := 0; row < b.Width; row++ {
		for col := 0; col < b.Width; col++ {
			// check if its a live cell, then apply rules for next generation

			// live cell
			if b.Cell[row][col] {
				aliveNeighbors := checkAliveNeighbors(b.Cell, row, col, b.Height, b.Width)

				// if less than 2 alive neighbors or more than 3 alive neighbors, cell dies
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					b.Cell[row][col] = false // cell dies
				}
			} else {
				// dead cell
				aliveNeighbors := checkAliveNeighbors(b.Cell, row, col, b.Height, b.Width)

				if aliveNeighbors == 3 {
					b.Cell[row][col] = true
				}
			}
		}
	}

}

func checkAliveNeighbors(cell [][]bool, row int, col int, height int, width int) int {

	// directions : [ up/down, left/right ] -1 = left, 1 = right, 1 = down, -1 = up
	directions := [][]int{
		{-1, -1}, // top left
		{-1, 0},  // top center
		{-1, 1},  // top right
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // bottom left
		{1, 0},   // bottom center
		{1, 1},   // bottom left
	}

	var aliveNeighbors int // tally the total neighbors
	// loop through directions to check all neighbors
	for _, d := range directions {
		// if any value is negative, we are looking off the board, stop looking
		newRowIndex := row + d[0]
		newColIndex := col + d[1]
		if newRowIndex < 0 || newRowIndex > width-1 || newColIndex < 0 || newColIndex > height-1 {
			continue
		}
		// fmt.Printf("newRowIndex %d ", newRowIndex)
		// fmt.Printf("newColIndex %d ", newColIndex)
		// fmt.Println("checking neighbor at:", row, col)
		neighborIsAlive := cell[row+d[0]][col+d[1]]
		// fmt.Println("neighbor is alive:", neighborIsAlive)

		if neighborIsAlive {
			aliveNeighbors++
		}
	}
	// fmt.Println("Total alive neighbors:", aliveNeighbors)

	return aliveNeighbors
}
