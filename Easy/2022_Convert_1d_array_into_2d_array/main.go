// https://leetcode.com/problems/convert-1d-array-into-2d-array/
package main

func construct2DArray(original []int, m int, n int) (mat [][]int) {
	if len(original) == m*n {
		mat = make([][]int, m)
		for i := 0; i < m; i++ {
			start := i * n
			mat[i] = original[start : start+n]
		}
	}
	return mat
}

func main() {}
