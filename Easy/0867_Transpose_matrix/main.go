// https://leetcode.com/problems/transpose-matrix/
package main

func transpose(matrix [][]int) (ret [][]int) {
	ret = make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		ret[i] = make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			ret[i][j] = matrix[j][i]
		}
	}
	return ret
}

func main() {}
