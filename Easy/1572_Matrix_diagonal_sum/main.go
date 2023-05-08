// https://leetcode.com/problems/matrix-diagonal-sum/
package main

func diagonalSum(mat [][]int) (ret int) {
	n := len(mat)
	for i := 0; i < n/2; i++ {
		ret += mat[i][i] + mat[n-i-1][n-i-1] + mat[i][n-i-1] + mat[n-i-1][i]
	}
	if n%2 == 1 {
		ret += mat[n/2][n/2]
	}
	return ret
}

func main() {}
