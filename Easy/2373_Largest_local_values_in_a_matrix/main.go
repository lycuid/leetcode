// https://leetcode.com/problems/largest-local-values-in-a-matrix/
package main

func largestLocal(grid [][]int) [][]int {
	ret := make([][]int, len(grid)-2)
	for i := range ret {
		ret[i] = make([]int, len(grid[i])-2)
	}
	for i := range ret {
		for j := range ret[i] {
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if num := grid[i+1+di][j+1+dj]; num > ret[i][j] {
						ret[i][j] = num
					}
				}
			}
		}
	}
	return ret
}

func main() {}
