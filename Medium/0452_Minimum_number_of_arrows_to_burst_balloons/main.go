// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
	cache := [][]int{points[0]}
	for i := 1; i < len(points); i++ {
		if last := len(cache) - 1; points[i][0] <= cache[last][1] {
			if cache[last][0] = points[i][0]; points[i][1] <= cache[last][1] {
				cache[last][1] = points[i][1]
			}
		} else {
			cache = append(cache, points[i])
		}
	}
	return len(cache)
}

func main() {}
