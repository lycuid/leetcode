// https://leetcode.com/problems/non-overlapping-intervals/
package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) (count int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	for i, j := 0, 1; j < len(intervals); j++ {
		if intervals[j][0] < intervals[i][1] {
			count++
		}
		if intervals[j][0] >= intervals[i][1] || intervals[j][1] < intervals[i][1] {
			i = j
		}
	}
	return count
}

func main() {}
