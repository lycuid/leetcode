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
	n := len(height)
	greater, water := make([]int, n), make([]int, n)
	greater[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		greater[i] = Max(height[i], greater[i+1])
	}
	for i := 1; i < n-1; i++ {
		water[i] = Max(0, Min(height[i-1], greater[i+1])-height[i]+water[i-1])
		ret += water[i]
	}
	return ret
}

func main() {}
