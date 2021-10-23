// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMin(xs []int) int {
	left, right := 0, len(xs)-1
	for (right - left) > 1 {
		if xs[left] > xs[right] {
			left += (right - left) / 2
		} else {
			right = left
			left /= 2
		}
	}
	return min(xs[left], xs[right])
}

func main() {}
