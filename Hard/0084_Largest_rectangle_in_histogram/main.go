// https://leetcode.com/problems/largest-rectangle-in-histogram/
package main

func largestRectangleArea(heights []int) (ret int) {
	n, cursor := len(heights), 0
	stack, previous, next := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		previous[i], next[i] = -1, n
	}
	for i := range heights {
		for ; cursor > 0 && heights[i] < heights[stack[cursor-1]]; cursor-- {
			next[stack[cursor-1]] = i
		}
		if cursor > 0 {
			previous[i] = stack[cursor-1]
		}
		stack[cursor], cursor = i, cursor+1
	}
	for i := range heights {
		if n := (next[i] - previous[i] - 1) * heights[i]; ret < n {
			ret = n
		}
	}
	return ret
}

func main() {}
