// https://leetcode.com/problems/range-sum-query-2d-immutable/
package main

type NumMatrix struct {
	inner [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	for i := range matrix {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := range matrix[i] {
			matrix[i][j] += matrix[i-1][j]
		}
	}
	return NumMatrix{matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (ret int) {
	ret = this.inner[row2][col2]
	if row1 > 0 {
		ret -= this.inner[row1-1][col2]
	}
	if col1 > 0 {
		ret -= this.inner[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		ret += this.inner[row1-1][col1-1]
	}
	return ret
}

func main() {}
