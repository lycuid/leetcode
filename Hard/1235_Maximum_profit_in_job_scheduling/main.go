// https://leetcode.com/problems/maximum-profit-in-job-scheduling/
package main

import "sort"

type Tuple struct{ start, end, value int }

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	cache, n := make([]Tuple, len(profit)+1), len(profit)
	for i := range profit {
		cache[i+1] = Tuple{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(cache, func(i, j int) bool { return cache[i].end < cache[j].end })
	for i := 1; i < len(cache); i++ {
		value, l, r := cache[i-1].value, 0, i-1
		for l < r {
			if mid := (l + r) / 2; cache[mid].end <= cache[i].start {
				l = mid + 1
			} else {
				r = mid
			}
		}
		for l > 0 && cache[l].end > cache[i].start {
			l--
		}
		if cache[i].value += cache[l].value; value > cache[i].value {
			cache[i].value = value
		}
	}
	return cache[n].value
}

func main() {}
