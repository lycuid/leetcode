// https://leetcode.com/problems/max-area-of-island/
package main

func F(i, j int, grid *[][]int) (ret int) {
	if i >= 0 && j >= 0 && i < len(*grid) && j < len((*grid)[0]) && (*grid)[i][j] != 0 {
		(*grid)[i][j] = 0
		ret = 1 + F(i-1, j, grid) + F(i+1, j, grid) + F(i, j-1, grid) + F(i, j+1, grid)
	}
	return ret
}

func maxAreaOfIsland(grid [][]int) (ret int) {
	for i := range grid {
		for j := range grid[i] {
			if tmp := F(i, j, &grid); tmp > ret {
				ret = tmp
			}
		}
	}
	return ret
}

func main() {}
