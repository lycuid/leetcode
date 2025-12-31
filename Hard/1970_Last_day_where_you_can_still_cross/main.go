// https://leetcode.com/problems/last-day-where-you-can-still-cross/
package main

const (
	Water = 1 << iota
	North = 1 << iota
	South = 1 << iota
)

func traverse(grid [][]int, y, x int) (ret bool) {
	if !(y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])) {
		return false
	}
	grid[y][x] = 0
	if y > 0 && grid[y-1][x] != Water {
		grid[y][x] |= grid[y-1][x]
	}
	if y < len(grid)-1 && grid[y+1][x] != Water {
		grid[y][x] |= grid[y+1][x]
	}
	if x > 0 && grid[y][x-1] != Water {
		grid[y][x] |= grid[y][x-1]
	}
	if x < len(grid[0])-1 && grid[y][x+1] != Water {
		grid[y][x] |= grid[y][x+1]
	}
	if grid[y][x] == North|South {
		return true
	}
	if y > 0 && grid[y-1][x] != grid[y][x] && grid[y-1][x] != Water {
		if traverse(grid, y-1, x) {
			return true
		}
	}
	if y < len(grid)-1 && grid[y+1][x] != grid[y][x] && grid[y+1][x] != Water {
		if traverse(grid, y+1, x) {
			return true
		}
	}
	if x > 0 && grid[y][x-1] != grid[y][x] && grid[y][x-1] != Water {
		if traverse(grid, y, x-1) {
			return true
		}
	}
	if x < len(grid[0])-1 && grid[y][x+1] != grid[y][x] && grid[y][x+1] != Water {
		if traverse(grid, y, x+1) {
			return true
		}
	}
	return false
}

func latestDayToCross(row, col int, cells [][]int) int {
	grid := make([][]int, row+2)
	for i := range grid {
		grid[i] = make([]int, col)
	}
	for i := range grid[0] {
		grid[0][i] = North
	}
	for i := range grid[row+1] {
		grid[row+1][i] = South
	}

	for _, cell := range cells {
		grid[cell[0]][cell[1]-1] |= Water
	}

	for i := len(cells) - 1; i >= 0; i-- {
		y, x := cells[i][0], cells[i][1]-1
		if grid[y][x] = 0; traverse(grid, y, x) {
			return i
		}
	}

	return 0
}

func main() {}
