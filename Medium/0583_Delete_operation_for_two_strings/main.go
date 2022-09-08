// https://leetcode.com/problems/delete-operation-for-two-strings/
package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			matrix[i][j] = Max(matrix[i-1][j], matrix[i][j-1])
			if word1[i-1] == word2[j-1] {
				matrix[i][j] = Max(matrix[i][j], matrix[i-1][j-1]+1)
			}
		}
	}
	return n + m - (2 * matrix[n][m])
}

func main() {}
