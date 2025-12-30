// https://leetcode.com/problems/magic-squares-in-grid/
package main

func isMagic(square [3][3]int) bool {
	var found [10]bool
	for _, row := range square {
		for _, num := range row {
			if num <= 9 {
				found[num] = true
			}
		}
	}
	unique := true
	for i := 1; unique && i < len(found); i++ {
		unique = unique && found[i]
	}

	sum := square[0][0] + square[0][1] + square[0][2]
	return unique &&
		square[1][0]+square[1][1]+square[1][2] == sum &&
		square[2][0]+square[2][1]+square[2][2] == sum &&
		square[0][0]+square[1][0]+square[2][0] == sum &&
		square[0][1]+square[1][1]+square[2][1] == sum &&
		square[0][2]+square[1][2]+square[2][2] == sum &&
		square[0][0]+square[1][1]+square[2][2] == sum &&
		square[0][2]+square[1][1]+square[2][0] == sum
}

func numMagicSquaresInside(grid [][]int) (count int) {
	var square [3][3]int
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[i])-3; j++ {
			square[0] = [3]int(grid[i][j : j+3])
			square[1] = [3]int(grid[i+1][j : j+3])
			square[2] = [3]int(grid[i+2][j : j+3])
			if isMagic(square) {
				count++
			}
		}
	}
	return count
}

func main() {}
