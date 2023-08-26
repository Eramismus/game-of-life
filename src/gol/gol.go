package gol

import (
	"math/rand"
	"time"
)

type Position struct {
	X int
	Y int
}

type Grid struct {
	Grid map[Position]bool
}

type JsonGrid struct {
	Grid [][]bool `json:"grid"`
}

func DefineGrid(width int, height int) (Grid, JsonGrid) {
	grid := Grid{Grid: make(map[Position]bool)}
	json_grid := JsonGrid{Grid: make([][]bool, width)}
	for i := 0; i < width; i++ {
		json_grid.Grid[i] = make([]bool, height)
		for j := 0; j < height; j++ {
			grid.Grid[Position{X: i, Y: j}] = false
		}
	}
	return grid, json_grid
}

func SeedNodes(grid *Grid, probability int) {
	for pos := range grid.Grid {
		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		random := generator.Intn(100)
		status := (random > probability)
		grid.Grid[pos] = status
	}
}

func FindNeighbours(grid Grid) map[Position][]Position {
	neighbours := make(map[Position][]Position)
	neighbour_range := []int{-1, 0, 1}

	for pos := range grid.Grid {
		X := pos.X
		Y := pos.Y

		var neighbour_X int
		var neighbour_Y int
		var neighbour_pos Position
		ns := make([]Position, 0)
		for _, offset_X := range neighbour_range {
			for _, offset_Y := range neighbour_range {
				neighbour_X = X + offset_X
				neighbour_Y = Y + offset_Y
				neighbour_pos = Position{X: neighbour_X, Y: neighbour_Y}

				if !((neighbour_X == pos.X) && (neighbour_Y == pos.Y)) {
					if _, ok := grid.Grid[neighbour_pos]; ok {
						ns = append(ns, neighbour_pos)
					}
				}
			}
		}
		neighbours[pos] = ns
	}
	return neighbours
}

func UpdateGrid(grid *Grid, neighbours map[Position][]Position, json_grid *JsonGrid) {

	for position := range grid.Grid {
		pos_ns := neighbours[position]

		trues := 0
		for _, n := range pos_ns {
			if grid.Grid[n] {
				trues++
			}
		}

		if grid.Grid[position] {
			if trues >= 2 && trues <= 3 {
				grid.Grid[position] = true
			} else {
				grid.Grid[position] = false
			}
		} else {
			if trues == 3 {
				grid.Grid[position] = true
			} else {
				grid.Grid[position] = false
			}
		}
		json_grid.Grid[position.X][position.Y] = grid.Grid[position]
	}
}
