// https://leetcode.com/problems/count-sub-islands/
package main

func removeIsland(i, j int, grid [][]int) {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] != 0 {
		grid[i][j] = 0
		removeIsland(i+1, j, grid)
		removeIsland(i-1, j, grid)
		removeIsland(i, j+1, grid)
		removeIsland(i, j-1, grid)
	}
}

func countSubIslands(grid1 [][]int, grid2 [][]int) (count int) {
	for i := range grid1 {
		for j := range grid1[i] {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				removeIsland(i, j, grid2)
			}
		}
	}
	for i := range grid2 {
		for j := range grid2[i] {
			if grid2[i][j] == 1 {
				removeIsland(i, j, grid2)
				count++
			}
		}
	}
	return count
}

func main() {}
