// https://leetcode.com/problems/insert-interval/
package main

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	for i := len(intervals) - 1; i > 0 && intervals[i][0] < intervals[i-1][0]; i-- {
		intervals[i], intervals[i-1] = intervals[i-1], intervals[i]
	}
	var i int
	for j := 1; j < len(intervals); j++ {
		if intervals[i][1] >= intervals[j][0] {
			if intervals[i][1] < intervals[j][1] {
				intervals[i][1] = intervals[j][1]
			}
		} else {
			intervals[i+1], i = intervals[j], i+1
		}
	}
	return intervals[:i+1]
}

func main() {}
