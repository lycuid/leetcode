// https://leetcode.com/problems/number-of-enclaves/
package main

func Aux(y, x int, grid [][]int) (ret int) {
	if grid[y][x] == 1 {
		grid[y][x], ret = 0, 1
		if y-1 >= 0 {
			ret += Aux(y-1, x, grid)
		}
		if y+1 < len(grid) {
			ret += Aux(y+1, x, grid)
		}
		if x-1 >= 0 {
			ret += Aux(y, x-1, grid)
		}
		if x+1 < len(grid[y]) {
			ret += Aux(y, x+1, grid)
		}
	}
	return ret
}

func numEnclaves(grid [][]int) (ret int) {
	for i := range grid {
		for j := range grid[i] {
			ret += grid[i][j]
		}
	}
	for _, i := range []int{0, len(grid) - 1} {
		for j := 0; j < len(grid[i]); j++ {
			ret -= Aux(i, j, grid)
		}
	}
	for i := 1; i < len(grid)-1; i++ {
		for _, j := range []int{0, len(grid[i]) - 1} {
			ret -= Aux(i, j, grid)
		}
	}
	return ret
}

func main() {}
