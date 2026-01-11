// https://leetcode.com/problems/maximal-rectangle/
package main

func maximalRectangle(matrix [][]byte) (area int) {
	if len(matrix) > 0 {
		for j := range matrix[0] {
			matrix[0][j] -= '0'
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = (matrix[i-1][j] + 1) * (matrix[i][j] - '0')
			}
		}
		s := make([]int, 1, len(matrix[0])+1)
		s[0] = -1
		for i := range matrix {
			for j := range matrix[i] {
				for ; len(s) > 1 && matrix[i][s[len(s)-1]] >= matrix[i][j]; s = s[:len(s)-1] {
					area = max(area, int(matrix[i][s[len(s)-1]])*(j-s[len(s)-2]-1))
				}
				s = append(s, j)
			}
			for ; len(s) > 1 && matrix[i][s[len(s)-1]] >= byte(0); s = s[:len(s)-1] {
				area = max(area, int(matrix[i][s[len(s)-1]])*(len(matrix[i])-s[len(s)-2]-1))
			}
			s = s[:1]
		}
	}
	return area
}

func main() {}
