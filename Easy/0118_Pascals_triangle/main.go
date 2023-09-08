// https://leetcode.com/problems/pascals-triangle/
package main

func generate(num int) [][]int {
	rows := make([][]int, num)
	for i := 0; i < num; i++ {
		rows[i] = make([]int, i+1)
		rows[i][0], rows[i][i] = 1, 1
		for j := 1; j < i; j++ {
			rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
		}
	}
	return rows
}

func main() {}
