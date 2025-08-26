// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/
package main

func areaOfMaxDiagonal(dimensions [][]int) int {
	var index, diag int
	for i, rect := range dimensions {
		if d := rect[0]*rect[0] + rect[1]*rect[1]; d > diag || d == diag && dimensions[index][0]*dimensions[index][1] < rect[0]*rect[1] {
			index, diag = i, d
		}
	}
	return dimensions[index][0] * dimensions[index][1]
}

func main() {}
