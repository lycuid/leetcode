// https://leetcode.com/problems/rotate-image/
package main

func rotate(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i <= n/2; i++ {
		for j := i; j <= n-i-1; j++ {
			matrix[i][j], matrix[n-j][i], matrix[n-i][n-j], matrix[j][n-i] = matrix[n-j][i], matrix[n-i][n-j], matrix[j][n-i], matrix[i][j]
		}
	}
}

func main() {}
