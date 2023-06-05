// https://leetcode.com/problems/check-if-it-is-a-straight-line/
package main

func checkStraightLine(coords [][]int) bool {
	y, x := coords[1][1]-coords[0][1], coords[1][0]-coords[0][0]
	for i := 1; i < len(coords); i++ {
		dy, dx := coords[i][1]-coords[i-1][1], coords[i][0]-coords[i-1][0]
		if dy*x != dx*y {
			return false
		}
	}
	return true
}

func main() {}
