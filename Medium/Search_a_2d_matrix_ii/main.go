// https://leetcode.com/problems/search-a-2d-matrix-ii/
package main

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

func main() {}
