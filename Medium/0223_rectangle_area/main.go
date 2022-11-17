// https://leetcode.com/problems/rectangle-area/
package main

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	return ((ay2 - ay1) * (ax2 - ax1)) +
		((by2 - by1) * (bx2 - bx1)) -
		(Max(Min(ay2, by2)-Max(ay1, by1), 0) * Max(Min(ax2, bx2)-Max(ax1, bx1), 0))
}

func main() {}
