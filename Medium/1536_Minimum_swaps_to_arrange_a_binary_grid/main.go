// https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/
package main

func minSwaps(grid [][]int) (count int) {
	n := len(grid)
	cache := make([]int, n)
	for i := range cache {
		for j := n - 1; j >= 0 && grid[i][j] == 0; j-- {
			cache[i]++
		}
	}
	for i := 0; i < n; i++ {
		var j int
		for j = i; j < n; j++ {
			if cache[j] >= n-1-i {
				for ; j > i; j-- {
					cache[j], cache[j-1], count = cache[j-1], cache[j], count+1
				}
				break
			}
		}
		if j == n {
			return -1
		}
	}
	return count
}

func main() {}
