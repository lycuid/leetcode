// https://leetcode.com/problems/special-positions-in-a-binary-matrix/
package main

func numSpecial(mat [][]int) (count int) {
	rows, cols := make([]int, len(mat)), make([]int, len(mat[0]))
	for i := range mat {
		for j := range mat[i] {
			rows[i], cols[j] = rows[i]+mat[i][j], cols[j]+mat[i][j]
		}
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				count++
			}
		}
	}
	return count
}

func main() {}
