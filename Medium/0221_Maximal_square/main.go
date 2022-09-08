// https://leetcode.com/problems/maximal-square/
package main

func Cmp(pred func(int, int) bool, arr ...int) int {
	if len(arr) == 0 {
		return 0
	}
	ret := arr[0]
	for _, val := range arr {
		if pred(val, ret) {
			ret = val
		}
	}
	return ret
}

func maximalSquare(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	m1 := make([]int, n*m)
	I := func(x, y int) int {
		return (m * x) + y
	}

	LT := func(x, y int) bool { return x < y }
	GT := func(x, y int) bool { return x > y }
	side := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				so_far := 0
				if i > 0 && j > 0 {
					so_far = Cmp(LT, m1[I(i-1, j)], m1[I(i, j-1)], m1[I(i-1, j-1)])
				}
				m1[I(i, j)] = so_far + 1
				side = Cmp(GT, side, m1[I(i, j)])
			}
		}
	}
	return side * side
}

func main() {}
