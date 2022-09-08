// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/
package main

func numSubmatrixSumTarget(grid [][]int, target int) (ret int) {
	rows, cols := len(grid), len(grid[0])
	sum := make([][]int, rows)
	for i := range sum {
		sum[i] = make([]int, cols)
		for j := range sum[i] {
			sum[i][j] = grid[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	MatrixSum := func(x1, y1, x2, y2 int) (ans int) {
		ans = sum[x2][y2]
		if x1 > 0 {
			ans -= sum[x1-1][y2]
		}
		if y1 > 0 {
			ans -= sum[x2][y1-1]
		}
		if x1 > 0 && y1 > 0 {
			ans += sum[x1-1][y1-1]
		}
		return ans
	}
	for x1 := 0; x1 < rows; x1++ {
		for x2 := x1; x2 < rows; x2++ {
			for y1 := 0; y1 < cols; y1++ {
				for y2 := y1; y2 < cols; y2++ {
					if MatrixSum(x1, y1, x2, y2) == target {
						ret++
					}
				}
			}
		}
	}
	return ret
}

func main() {}
