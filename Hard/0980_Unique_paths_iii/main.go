// https://leetcode.com/problems/unique-paths-iii/
package main

func Aux(y, x, visited int, grid [][]int) (ret int) {
	val, m, n := grid[y][x], len(grid), len(grid[0])
	if val == 2 && visited == m*n-1 {
		return 1
	}
	grid[y][x], visited = -1, visited+1
	for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		dy, dx := y+d[0], x+d[1]
		if dy >= 0 && dy < m && dx >= 0 && dx < n && grid[dy][dx] != -1 {
			ret += Aux(dy, dx, visited, grid)
		}
	}
	grid[y][x], visited = val, visited-1
	return ret
}

func uniquePathsIII(grid [][]int) int {
	sy, sx, visited := 0, 0, 0
	for y := range grid {
		for x := range grid[y] {
			switch grid[y][x] {
			case -1:
				visited++
			case 1:
				sy, sx = y, x
			}
		}
	}
	return Aux(sy, sx, visited, grid)
}

func main() {}
