// https://leetcode.com/problems/count-square-submatrices-with-all-ones/
package main

func countSquares(matrix [][]int) (count int) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 1 && i > 0 && j > 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1])
			}
			count += matrix[i][j]
		}
	}
	return count
}

func main() {}
