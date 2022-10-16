// https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/
package main

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	cache := make([][]int, d+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	cache[1][1] = jobDifficulty[0]
	for i := 1; i <= n; i++ {
		if cache[1][i] = cache[1][i-1]; cache[1][i] < jobDifficulty[i-1] {
			cache[1][i] = jobDifficulty[i-1]
		}
	}
	for i := 2; i <= d; i++ {
		for j := i; j <= n; j++ {
			cache[i][j] = 1e9 + 7
			for k, m := j-1, jobDifficulty[j-1]; k >= i-1; k-- {
				if cache[i][j] > cache[i-1][k]+m {
					cache[i][j] = cache[i-1][k] + m
				}
				if m < jobDifficulty[k-1] {
					m = jobDifficulty[k-1]
				}
			}
		}
	}
	return cache[d][n]
}

func main() {}
