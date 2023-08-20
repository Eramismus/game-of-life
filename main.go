package main

import (
	// "math/rand"
	"fmt"

	. "github.com/eramismus/game-of-life/src/gol"
	. "github.com/eramismus/game-of-life/src/graphics"
)

func main() {
	fmt.Println("Starting the game")
	StartWebSocketServer(10000, "/endpoint")

	game_grid := DefineGrid(5, 5)
	fmt.Println("Nodes:", game_grid)

	// Seed the nodes
	positions := make([]Position, 10)
	positions[0] = Position{X: 0, Y: 1}
	positions[1] = Position{X: 1, Y: 0}
	positions[2] = Position{X: 1, Y: 1}
	positions[3] = Position{X: 1, Y: 2}

	SeedNodes(&game_grid, positions)
	fmt.Println("Seeded Nodes:", game_grid)

	neighbour_map := FindNeighbours(game_grid)
	fmt.Println("Neighbour map:", neighbour_map)

	for i := 0; i < 20; i++ {
		UpdateGrid(&game_grid, neighbour_map)
		fmt.Println("Nodes:", game_grid)
	}
}
