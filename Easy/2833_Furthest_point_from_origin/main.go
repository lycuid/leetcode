// https://leetcode.com/problems/furthest-point-from-origin/
package main

func furthestDistanceFromOrigin(moves string) int {
	var c1, c2 int
	for _, move := range moves {
		switch move {
		case 'R':
			c1++
		case 'L':
			c1--
		default:
			c2++
		}
	}
	if c1 < 0 {
		c1 = -c1
	}
	return c1 + c2
}

func main() {}
