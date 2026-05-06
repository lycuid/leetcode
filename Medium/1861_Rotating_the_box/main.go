// https://leetcode.com/problems/rotating-the-box/
package main

func rotateTheBox(boxGrid [][]byte) [][]byte {
	res := make([][]byte, len(boxGrid[0]))
	for i := range res {
		res[i] = make([]byte, len(boxGrid))
	}
	for i := range boxGrid {
		drop, _i := len(boxGrid[i])-1, len(boxGrid)-1-i
		for j := len(boxGrid[i]) - 1; j >= 0; j-- {
			switch boxGrid[i][j] {
			case '.':
				res[j][_i] = '.'
			case '*':
				res[j][_i], drop = '*', j-1
			case '#':
				res[j][_i], res[drop][_i], drop = '.', '#', drop-1
			}
		}
	}
	return res
}

func main() {}
