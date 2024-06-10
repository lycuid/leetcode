// https://leetcode.com/problems/height-checker/
package main

import "sort"

func heightChecker(heights []int) (count int) {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	for i, num := range expected {
		if num != heights[i] {
			count++
		}
	}
	return count
}

func main() {}
