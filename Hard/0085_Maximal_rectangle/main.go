// https://leetcode.com/problems/maximal-rectangle/
package main

func maximalRectangle(matrix [][]byte) (ret int) {
	cache := make([][]int, len(matrix))
	for i := range cache {
		cache[i] = make([]int, len(matrix[i]))
		if len(cache[i]) > 0 {
			cache[i][0] = int(matrix[i][0] - '0')
		}
		for j := 1; j < len(cache[i]); j++ {
			if matrix[i][j] != '0' {
				cache[i][j] = cache[i][j-1] + int(matrix[i][j]-'0')
			}
		}
	}
	for i := range cache {
		for j := range cache[i] {
			for k, min := 0, cache[i][j]; k <= i; k++ {
				if min > cache[i-k][j] {
					min = cache[i-k][j]
				}
				if n := (k + 1) * min; n > ret {
					ret = n
				}
			}
		}
	}
	return ret
}

func main() {}
