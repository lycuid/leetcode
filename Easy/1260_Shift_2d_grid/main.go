// https://leetcode.com/problems/shift-2d-grid/
package main

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	mxn := m * n
	if k = mxn - (k % mxn); k > 0 {
		swap := func(i, j int) {
			xi, yi := i/n, i%n
			xj, yj := j/n, j%n
			grid[xi][yi], grid[xj][yj] = grid[xj][yj], grid[xi][yi]
		}
		for i := 0; i < k/2; i++ {
			swap(i, k-1-i)
		}
		for i := 0; i < (mxn-k)/2; i++ {
			swap(k+i, mxn-1-i)
		}
		for i := 0; i < mxn/2; i++ {
			swap(i, mxn-1-i)
		}
	}
	return grid
}

func main() {}
