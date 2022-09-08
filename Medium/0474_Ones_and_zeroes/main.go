// https://leetcode.com/problems/ones-and-zeroes/
package main

type Point struct {
	x, y int
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMaxForm(strs []string, x int, y int) (count int) {
	points := make([]Point, len(strs))
	for i, str := range strs {
		points[i] = Point{}
		for _, ch := range str {
			if ch == '0' {
				points[i].x++
			}
			if ch == '1' {
				points[i].y++
			}
		}
	}
	cache := make([][]int, x+1)
	for i := range cache {
		cache[i] = make([]int, y+1)
	}
	for _, point := range points {
		px, py := point.x, point.y
		for i := x; i >= px; i-- {
			for j := y; j >= py; j-- {
				cache[i][j] = Max(cache[i][j], 1+cache[i-px][j-py])
			}
		}
	}
	return cache[x][y]
}

func main() {}
