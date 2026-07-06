// https://leetcode.com/problems/remove-covered-intervals/
package main

import "sort"

func removeCoveredIntervals(intervals [][]int) (res int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var high int
	for _, i := range intervals {
		if i[1] > high {
			high = i[1]
		} else {
			res++
		}
	}
	return len(intervals) - res
}

func main() {}
