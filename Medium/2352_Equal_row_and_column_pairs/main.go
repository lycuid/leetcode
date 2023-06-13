// https://leetcode.com/problems/equal-row-and-column-pairs/
package main

func equalPairs(grid [][]int) (ret int) {
	n := len(grid)
	for i := 0; i < n; i++ {
	LOOP:
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if grid[i][k] != grid[k][j] {
					continue LOOP
				}
			}
			ret++
		}
	}
	return ret
}

func main() {}
