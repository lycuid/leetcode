// https://leetcode.com/problems/path-crossing/
package main

func isPathCrossing(path string) bool {
	type Point struct{ y, x int }

	cache := make(map[Point]bool)
	cache[Point{0, 0}] = true

	var point Point
	for i := range path {
		switch path[i] {
		case 'N':
			point.y++
		case 'E':
			point.x++
		case 'S':
			point.y--
		case 'W':
			point.x--
		}
		if _, found := cache[point]; found {
			return true
		}
		cache[point] = true
	}
	return false
}

func main() {}
