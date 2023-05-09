// https://leetcode.com/problems/spiral-matrix/
package main

func spiralOrder(matrix [][]int) []int {
	ret := make([]int, len(matrix)*len(matrix[0]))
	y0, x0, y1, x1 := 0, 0, len(matrix)-1, len(matrix[0])-1
	for c := 0; c < len(ret); {
		for i := x0; i <= x1; i++ {
			ret[c], c = matrix[y0][i], c+1
		}
		for i := y0 + 1; i < y1; i++ {
			ret[c], c = matrix[i][x1], c+1
		}
		if y1 > y0 {
			for i := x1; i >= x0; i-- {
				ret[c], c = matrix[y1][i], c+1
			}
		}
		if x1 > x0 {
			for i := y1 - 1; i > y0; i-- {
				ret[c], c = matrix[i][x0], c+1
			}
		}
		y0, x0, y1, x1 = y0+1, x0+1, y1-1, x1-1
	}
	return ret
}

func main() {}
