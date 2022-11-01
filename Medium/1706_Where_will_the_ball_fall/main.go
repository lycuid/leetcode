// https://leetcode.com/problems/where-will-the-ball-fall/
package main

func findBall(grid [][]int) (ret []int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		ret = append(ret, -1)
	}
OUTTER:
	for c, j := 0, 0; c < n; c, j = c+1, c+1 {
		for i, k := 0, 0; i < m; i, j = i+1, k {
			if k = j + grid[i][j]; k < 0 || k >= n || grid[i][j]+grid[i][k] == 0 {
				continue OUTTER
			}
		}
		ret[c] = j
	}
	return ret
}

func main() {}
