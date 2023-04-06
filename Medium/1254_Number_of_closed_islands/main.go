// https://leetcode.com/problems/number-of-closed-islands/
package main

func Aux(y, x int, grid [][]int) bool {
	if grid[y][x] == 1 {
		return true
	}
	grid[y][x] = 1
	var left, right, top, bottom bool
	if x-1 >= 0 {
		left = Aux(y, x-1, grid)
	}
	if x+1 < len(grid[0]) {
		right = Aux(y, x+1, grid)
	}
	if y-1 >= 0 {
		top = Aux(y-1, x, grid)
	}
	if y+1 < len(grid) {
		bottom = Aux(y+1, x, grid)
	}
	return left && right && top && bottom
}

func closedIsland(grid [][]int) (ret int) {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 && Aux(y, x, grid) {
				ret++
			}
		}
	}
	return ret
}

func main() {}
