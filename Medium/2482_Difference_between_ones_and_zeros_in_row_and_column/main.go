// https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/
package main

func onesMinusZeros(grid [][]int) [][]int {
	diff := make([][]int, len(grid))
	for i := range grid {
		diff[i] = make([]int, len(grid[i]))
	}
	rows, cols := make([][2]int, len(grid)), make([][2]int, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			rows[i][grid[i][j]]++
			cols[j][grid[i][j]]++
		}
	}
	for i := range grid {
		for j := range grid[i] {
			diff[i][j] = rows[i][1] + cols[j][1] - rows[i][0] - cols[j][0]
		}
	}
	return diff
}

func main() {}
