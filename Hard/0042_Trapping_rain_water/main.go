// https://leetcode.com/problems/trapping-rain-water/
package main

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func trap(height []int) (trapped int) {
	right := make([]int, len(height)+1)
	for i := len(right) - 2; i >= 0; i-- {
		right[i] = Max(height[i], right[i+1])
	}
	for i, left := 0, 0; i < len(height); i, left = i+1, Max(left, height[i]) {
		wall := right[i+1]
		if wall > left {
			wall = left
		}
		trapped += Max(0, wall-height[i])
	}
	return trapped
}

func main() {}
