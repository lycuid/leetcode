// https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/
package main

func countSubmatrices(grid [][]int, k int) (count int) {
	for i := range grid {
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] += grid[i][j-1]
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if i > 0 {
				grid[i][j] += grid[i-1][j]
			}
			if grid[i][j] <= k {
				count++
			}
		}
	}
	return count
}

func main() {}
