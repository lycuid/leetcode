// https://leetcode.com/problems/minimum-falling-path-sum/
package main

func Min(num int, nums ...int) int {
	for i := range nums {
		if nums[i] < num {
			num = nums[i]
		}
	}
	return num
}

func minFallingPathSum(matrix [][]int) int {
	cache, n := make([][]int, len(matrix)+1), len(matrix)
	for i := range cache {
		cache[i] = make([]int, len(matrix[0]))
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < len(cache[i]); j++ {
			if cache[i][j] = matrix[i-1][j] + cache[i-1][j]; j > 0 {
				cache[i][j] = Min(cache[i][j], matrix[i-1][j]+cache[i-1][j-1])
			}
			if j < len(cache[i])-1 {
				cache[i][j] = Min(cache[i][j], matrix[i-1][j]+cache[i-1][j+1])
			}
		}
	}
	return Min(cache[n][0], cache[n]...)
}

func main() {}
