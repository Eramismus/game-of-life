package gol

type Position struct {
	X int
	Y int
}

type Grid struct {
	grid map[Position]bool
}

func DefineGrid(width int, height int) Grid {
	grid := Grid{grid: make(map[Position]bool)}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			grid.grid[Position{X: i, Y: j}] = false
		}
	}
	return grid
}

func SeedNodes(grid *Grid, positions []Position) {
	for _, pos := range positions {
		grid.grid[pos] = true
	}
}

func FindNeighbours(grid Grid) map[Position][]Position {
	neighbours := make(map[Position][]Position)
	neighbour_range := []int{-1, 0, 1}

	for pos := range grid.grid {
		X := pos.X
		Y := pos.Y

		var neighbour_X int
		var neighbour_Y int
		var neighbour_pos Position
		ns := make([]Position, 0)
		for offset_X := range neighbour_range {
			for offset_Y := range neighbour_range {
				neighbour_X = X + offset_X
				neighbour_Y = Y + offset_Y
				neighbour_pos = Position{X: offset_X, Y: offset_Y}

				if neighbour_X != pos.X && neighbour_Y != pos.Y {
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

func UpdateGrid(grid *Grid, neighbours map[Position][]Position) {

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
