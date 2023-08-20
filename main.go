package main

import (
	"fmt"
	// "math/rand"
)

type Position struct {
	x int
	y int
}

type Grid struct {
	grid map[Position]bool
}

func main() {
	fmt.Println("Starting the game")
	game_grid := defineGrid(5, 5)
	fmt.Println("Nodes:", game_grid)

	// Seed the nodes
	positions := make([]Position, 10)
	positions[0] = Position{x: 0, y: 1}
	positions[1] = Position{x: 1, y: 0}
	positions[2] = Position{x: 1, y: 1}
	positions[3] = Position{x: 1, y: 2}
	// positions[4] = Position{x: 3, y: 2}
	// positions[5] = Position{x: 2, y: 2}
	// positions[6] = Position{x: 3, y: 4}
	// positions[7] = Position{x: 3, y: 5}
	// positions[8] = Position{x: 4, y: 5}
	// positions[9] = Position{x: 2, y: 1}

	seedNodes(&game_grid, positions)
	fmt.Println("Seeded Nodes:", game_grid)

	neighbour_map := findNeighbours(game_grid)
    fmt.Println("Neighbour map:", neighbour_map)

	for i := 0; i < 20; i++ {
		updateGrid(&game_grid, neighbour_map)
		fmt.Println("Nodes:", game_grid)
	}
}

func defineGrid(width int, height int) Grid {
	grid := Grid{grid: make(map[Position]bool)}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			grid.grid[Position{x: i, y: j}] = false
		}
	}
	return grid
}

func seedNodes(grid *Grid, positions []Position) {
	for _, pos := range positions {
		grid.grid[pos] = true
	}
}

func findNeighbours(grid Grid) map[Position][]Position {
	neighbours := make(map[Position][]Position)
	neighbour_range := []int{-1, 0, 1}

	for pos := range grid.grid {
		x := pos.x
		y := pos.y

		var neighbour_x int
		var neighbour_y int
		var neighbour_pos Position
        ns := make([]Position,0)
		for offset_x := range neighbour_range {
			for offset_y := range neighbour_range {
				neighbour_x = x + offset_x
				neighbour_y = y + offset_y
				neighbour_pos = Position{x: offset_x, y: offset_y}

				if neighbour_x != pos.x && neighbour_y != pos.y {
					if _, ok := grid.grid[neighbour_pos]; ok {
                        ns = append(ns, neighbour_pos)
					}
				}
			}
		}
        neighbours[pos] = ns
	}
	return neighbours
}

func updateGrid(grid *Grid, neighbours map[Position][]Position) {
    
    for position := range grid.grid {
        pos_ns := neighbours[position]
        
        trues := 0 
        for _, n := range pos_ns {
            if grid.grid[n] {
                trues++
            }
        }

        if grid.grid[position] {
            if trues >= 2 && trues <= 3 {
                grid.grid[position] = true
            } else {
                grid.grid[position] = false
            }
        } else {
            if trues == 3 {
                grid.grid[position] = true
            } else {
                grid.grid[position] = false
            }
        }
    }
}
