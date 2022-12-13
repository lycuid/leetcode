// https://leetcode.com/problems/minimum-falling-path-sum/
package main

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minFallingPathSum(matrix [][]int) (ret int) {
	for i, j := 0, 0; i < len(matrix); i++ {
		for j, ret = 0, 1e9+7; j < len(matrix[i]); j++ {
			if n := matrix[i][j]; i > 0 {
				if n += matrix[i-1][j]; j > 0 {
					n = Min(n, matrix[i][j]+matrix[i-1][j-1])
				}
				if j < len(matrix[i])-1 {
					n = Min(n, matrix[i][j]+matrix[i-1][j+1])
				}
				matrix[i][j] = n
			}
			ret = Min(ret, matrix[i][j])
		}
	}
	return ret
}

func main() {}
