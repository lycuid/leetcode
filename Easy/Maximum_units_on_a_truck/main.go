// https://leetcode.com/problems/maximum-units-on-a-truck/
package main

import "sort"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maximumUnits(boxes [][]int, n int) (ret int) {
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i][1] > boxes[j][1]
	})
	for i := 0; i < len(boxes) && n > 0; i++ {
		min := Min(boxes[i][0], n)
		ret += min * boxes[i][1]
		n -= min
	}
	return ret
}

func main() {}
