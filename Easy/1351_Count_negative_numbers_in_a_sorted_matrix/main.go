// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
package main

func countNegatives(grid [][]int) (ret int) {
	for i, r := 0, len(grid[0]); i < len(grid); i++ {
		l := 0
		for l < r {
			if mid := (l + r) / 2; grid[i][mid] < 0 {
				r = mid
			} else {
				l = mid + 1
			}
		}
		ret += len(grid[i]) - l
	}
	return ret
}

func main() {}
