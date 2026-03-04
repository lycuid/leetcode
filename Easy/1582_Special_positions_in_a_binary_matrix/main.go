// https://leetcode.com/problems/special-positions-in-a-binary-matrix/
package main

func numSpecial(mat [][]int) (count int) {
	var cache [2][]int
	cache[0] = make([]int, len(mat))
	cache[1] = make([]int, len(mat[0]))
	for i := range mat {
		for j := range mat[i] {
			cache[0][i] += mat[i][j]
			cache[1][j] += mat[i][j]
		}
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && cache[0][i] == 1 && cache[1][j] == 1 {
				count++
			}
		}
	}
	return count
}

func main() {}
