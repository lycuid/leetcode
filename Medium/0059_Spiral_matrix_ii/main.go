// https://leetcode.com/problems/spiral-matrix-ii/
package main

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	start, end, val := 0, n, 1
	for start < end {
		for i := start; i < end; i++ {
			ret[start][i] = val
			val++
		}
		for i := start + 1; i < end; i++ {
			ret[i][end-1] = val
			val++
		}
		for i := end - 2; i >= start; i-- {
			ret[end-1][i] = val
			val++
		}
		for i := end - 2; i > start; i-- {
			ret[i][start] = val
			val++
		}
		start++
		end--
	}
	return ret
}

func main() {}
