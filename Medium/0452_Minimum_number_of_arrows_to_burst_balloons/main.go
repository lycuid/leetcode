// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
package main

import "sort"

func findMinArrowShots(points [][]int) (i int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for j := 0; j < len(points); j++ {
		if points[j][0] <= points[i][1] {
			if points[i][1] > points[j][1] {
				points[i][1] = points[j][1]
			}
		} else {
			points[i+1], i = points[j], i+1
		}
	}
	return i + 1
}

func main() {}
