// https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/
package main

import "sort"

func maxWidthOfVerticalArea(points [][]int) (max int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for i := 1; i < len(points)-1; i++ {
		if n := points[i][0] - points[i-1][0]; n > max {
			max = n
		}
	}
	return max
}

func main() {}
