// https://leetcode.com/problems/two-best-non-overlapping-events/
package main

import "sort"

func maxTwoEvents(events [][]int) (maxValue int) {
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] == events[j][1] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})
	cache := make([]int, len(events)+1)
	for i, e := range events {
		cache[i+1] = max(cache[i], e[2])
		maxValue = max(maxValue, e[2])
		l, r := 0, i-1
		for l <= r {
			if mid := l + (r-l)/2; events[mid][1] < e[0] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if r >= 0 {
			maxValue = max(maxValue, e[2]+cache[r+1])
		}
	}
	return maxValue
}

func main() {}
