// https://leetcode.com/problems/container-with-most-water/
package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArea(height []int) (max_so_far int) {
	l, r := 0, len(height)-1
	for l < r {
		max_so_far = max(min(height[l], height[r])*(r-l), max_so_far)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max_so_far
}

func main() {}
