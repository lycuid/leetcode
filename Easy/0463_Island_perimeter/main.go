// https://leetcode.com/problems/island-perimeter/
package main

func islandPerimeter(grid [][]int) (perimeter int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				if i == 0 || grid[i-1][j] == 0 {
					perimeter++
				}
				if i == len(grid)-1 || grid[i+1][j] == 0 {
					perimeter++
				}
				if j == 0 || grid[i][j-1] == 0 {
					perimeter++
				}
				if j == len(grid[i])-1 || grid[i][j+1] == 0 {
					perimeter++
				}
			}
		}
	}
	return perimeter
}

func main() {}
