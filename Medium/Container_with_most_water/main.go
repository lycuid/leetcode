// https://leetcode.com/problems/container-with-most-water/

package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func maxArea(height []int) (maxSoFar int) {
	left := 0
	right := len(height) - 1
	for left < right {
		maxSoFar = max(min(height[left], height[right])*(right-left), maxSoFar)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxSoFar
}

func main() {}
