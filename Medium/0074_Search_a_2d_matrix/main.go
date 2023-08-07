// https://leetcode.com/problems/search-a-2d-matrix/
package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l < r {
		if mid := (l + r) / 2; matrix[mid/n][mid%n] > target {
			r = mid
		} else if matrix[mid/n][mid%n] < target {
			l = mid + 1
		} else {
			l, r = mid, mid
		}
	}
	return matrix[l/n][l%n] == target
}

func main() {}
