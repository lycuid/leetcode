// https://leetcode.com/problems/shift-2d-grid/
package main

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	l := m * n
	k = l - (k % l)
	xy := func(i int) (int, int) {
		return i / n, i % n
	}
	swap := func(i, j int) {
		xi, yi := xy(i)
		xj, yj := xy(j)
		grid[xi][yi], grid[xj][yj] = grid[xj][yj], grid[xi][yi]
	}
	if k%l == 0 {
		goto EXIT
	}
	for i := 0; i < k/2; i++ {
		swap(i, k-i-1)
	}
	for i := 0; i < (l-k)/2; i++ {
		swap(i+k, l-i-1)
	}
	for i := 0; i < l/2; i++ {
		swap(i, l-i-1)
	}
EXIT:
	return grid
}

func main() {}
