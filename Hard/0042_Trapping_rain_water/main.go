// https://leetcode.com/problems/trapping-rain-water/
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

func trap(height []int) (ret int) {
	n, post := len(height), make([]int, len(height)+1)
	for i := n - 1; i >= 0; i-- {
		post[i] = Max(height[i], post[i+1])
	}
	for i, prev := 0, 0; i < n; i, prev = i+1, Max(prev, height[i]) {
		ret += Max(0, Min(prev, post[i+1])-height[i])
	}
	return ret
}

func main() {}
