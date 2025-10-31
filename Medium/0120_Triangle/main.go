// https://leetcode.com/problems/triangle/
package main

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 1; i > 0; i-- {
		for j := 0; j < len(triangle[i])-1; j++ {
			triangle[i-1][j] += min(triangle[i][j], triangle[i][j+1])
		}
	}
	return triangle[0][0]
}

func main() {}
