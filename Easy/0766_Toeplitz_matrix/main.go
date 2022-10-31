// https://leetcode.com/problems/toeplitz-matrix/
package main

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix) && j-i < len(matrix[0]); j++ {
			if matrix[i][0] != matrix[j][j-i] {
				return false
			}
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		for j := i + 1; j-i < len(matrix) && j < len(matrix[0]); j++ {
			if matrix[0][i] != matrix[j-i][j] {
				return false
			}
		}
	}
	return true
}

func main() {}
