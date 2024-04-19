// https://leetcode.com/problems/number-of-islands/
package main

func Aux(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	if i > 0 && grid[i-1][j] == '1' {
		Aux(grid, i-1, j)
	}
	if i < len(grid)-1 && grid[i+1][j] == '1' {
		Aux(grid, i+1, j)
	}
	if j > 0 && grid[i][j-1] == '1' {
		Aux(grid, i, j-1)
	}
	if j < len(grid[i])-1 && grid[i][j+1] == '1' {
		Aux(grid, i, j+1)
	}
}

func numIslands(grid [][]byte) (count int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				Aux(grid, i, j)
			}
		}
	}
	return count
}

func main() {}
