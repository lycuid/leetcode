// https://leetcode.com/problems/spiral-matrix-ii/
package main

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	x0, y0, x1, y1 := 0, 0, n-1, n-1
	for num, m := 1, n*n; num <= m; x0, y0, x1, y1 = x0+1, y0+1, x1-1, y1-1 {
		for i := x0; i <= x1; i++ {
			ret[y0][i], num = num, num+1
		}
		for i := y0 + 1; i <= y1-1; i++ {
			ret[i][x1], num = num, num+1
		}
		if y1 > y0 {
			for i := x1; i >= x0; i-- {
				ret[y1][i], num = num, num+1
			}
		}
		if x1 > x0 {
			for i := y1 - 1; i >= y0+1; i-- {
				ret[i][x0], num = num, num+1
			}
		}
	}
	return ret
}

func main() {}
