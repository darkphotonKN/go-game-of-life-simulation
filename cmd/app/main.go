package main

import (
	"github.com/darkphotonKN/go-game-of-life-simulation/internal/app/game"
	"github.com/fatih/color"
)

func main() {
	color.Blue("Conways Game of Life Simulation")

	gameOfLife := game.NewBoard()

	gameOfLife.Simulation()
}
