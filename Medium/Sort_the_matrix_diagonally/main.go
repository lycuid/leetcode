// https://leetcode.com/problems/sort-the-matrix-diagonally/
package main

import "sort"

func SortDiagonal(sy, sx int, matrix [][]int) {
	var diagonal []int
	r, c := len(matrix), len(matrix[0])
	for y, x := sy, sx; y < r && x < c; y, x = y+1, x+1 {
		diagonal = append(diagonal, matrix[y][x])
	}
	sort.Ints(diagonal)
	for i := 0; i < len(diagonal); i, sy, sx = i+1, sy+1, sx+1 {
		matrix[sy][sx] = diagonal[i]
	}
}

func diagonalSort(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		SortDiagonal(i, 0, matrix)
	}
	for i := 1; i < n; i++ {
		SortDiagonal(0, i, matrix)
	}
	return matrix
}

func main() {}
