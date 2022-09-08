// https://leetcode.com/problems/triangle/
package main

import "math"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minimumTotal(triangle [][]int) int {
	res, n := math.MaxInt, len(triangle)
	for i := 1; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += Min(triangle[i-1][Max(j-1, 0)],
				triangle[i-1][Min(j, len(triangle[i-1])-1)])
		}
	}
	for i := 0; i < n; i++ {
		res = Min(res, triangle[n-1][i])
	}
	return res
}

func main() {}
