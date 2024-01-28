// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/
package main

func numSubmatrixSumTarget(matrix [][]int, target int) (count int) {
	cache := make([][]int, len(matrix)+1)
	for i := range cache {
		cache[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(cache); i++ {
		for j := 1; j < len(cache[i]); j++ {
			cache[i][j] = matrix[i-1][j-1] + cache[i-1][j] + cache[i][j-1] - cache[i-1][j-1]
		}
	}
	for x1 := 1; x1 < len(cache); x1++ {
		for y1 := 1; y1 < len(cache[x1]); y1++ {
			for x2 := x1; x2 < len(cache); x2++ {
				for y2 := y1; y2 < len(cache[x2]); y2++ {
					if cache[x2][y2]-cache[x1-1][y2]-cache[x2][y1-1]+cache[x1-1][y1-1] == target {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {}
