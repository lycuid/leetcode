// https://leetcode.com/problems/transpose-matrix/
package main

func transpose(matrix [][]int) [][]int {
	ret := make([][]int, len(matrix[0]))
	for i := range ret {
		ret[i] = make([]int, len(matrix))
	}
	for i := range matrix {
		for j := range matrix[i] {
			ret[j][i] = matrix[i][j]
		}
	}
	return ret
}

func main() {}
