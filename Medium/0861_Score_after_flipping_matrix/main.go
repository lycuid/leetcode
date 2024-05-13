// https://leetcode.com/problems/score-after-flipping-matrix/
package main

func matrixScore(grid [][]int) (total int) {
	if m := len(grid); m > 0 {
		if n := len(grid[0]); n > 0 {
			for j := 0; j < n; j++ {
				var count [2]int
				for i := 0; i < m; i++ {
					if grid[i][0] == 1 {
						count[grid[i][j]]++
					} else {
						count[1-grid[i][j]]++
					}
				}
				if count[0] < count[1] {
					count[0] = count[1]
				}
				total += (1 << (n - j - 1)) * count[0]
			}
		}
	}
	return total
}

func main() {}
