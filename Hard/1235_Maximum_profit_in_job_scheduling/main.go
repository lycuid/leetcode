// https://leetcode.com/problems/maximum-profit-in-job-scheduling/
package main

import "sort"

type Job struct{ start, end, profit int }

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(profit)
	job, cache := make([]Job, n), make([]int, n+1)
	for i := range job {
		job[i] = Job{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(job, func(i, j int) bool { return job[i].start < job[j].start })
	for i, l, r := n-1, 0, 0; i >= 0; i-- {
		cache[i], l, r = job[i].profit, i+1, n-1
		for l < r {
			if mid := (l + r) / 2; job[mid].start < job[i].end {
				l = mid + 1
			} else if job[mid].start > job[i].end {
				r = mid
			} else {
				break
			}
		}
		for ; l < n; l++ {
			if job[i].end <= job[l].start {
				cache[i] += cache[l]
				break
			}
		}
		if cache[i] < cache[i+1] {
			cache[i] = cache[i+1]
		}
	}
	return cache[0]
}

func main() {}
