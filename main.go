package main

import (
	// "math/rand"
	"fmt"

	. "github.com/eramismus/game-of-life/src/gol"
	. "github.com/eramismus/game-of-life/src/graphics"
)

func main() {
	fmt.Println("Starting the game")
	game_grid, json_grid := DefineGrid(100, 100)

	// Seed the nodes
	SeedNodes(&game_grid, 70)
	neighbour_map := FindNeighbours(game_grid)

	StartWebSocketServer(10000, "/endpoint", game_grid, neighbour_map, json_grid)
}
