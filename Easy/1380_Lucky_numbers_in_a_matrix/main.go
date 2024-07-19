// https://leetcode.com/problems/lucky-numbers-in-a-matrix/
package main

func luckyNumbers(matrix [][]int) (ret []int) {
	min_row := make([]int, len(matrix))
	for i, row := range matrix {
		for j, val := range row {
			if val < matrix[i][min_row[i]] {
				min_row[i] = j
			}
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		var max_col int
		for i := 0; i < len(matrix); i++ {
			max_col = max(max_col, matrix[i][j])
		}
		for i := 0; i < len(matrix); i++ {
			if max_col == matrix[i][min_row[i]] {
				ret = append(ret, matrix[i][min_row[i]])
			}
		}
	}
	return ret
}

func main() {}
