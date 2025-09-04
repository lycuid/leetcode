// https://leetcode.com/problems/find-closest-person/
package main

func findClosest(x int, y int, z int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	if dx, dy := abs(z-x), abs(z-y); dx < dy {
		return 1
	} else if dx > dy {
		return 2
	}
	return 0
}

func main() {}
