// https://leetcode.com/problems/data-stream-as-disjoint-intervals/
package main

import "sort"

type SummaryRanges struct{ inner []int }

func Constructor() SummaryRanges            { return SummaryRanges{make([]int, 0)} }
func (this *SummaryRanges) AddNum(item int) { this.inner = append(this.inner, item) }

func (this *SummaryRanges) GetIntervals() (intervals [][]int) {
	sort.Ints(this.inner)
	if len(this.inner) > 0 {
		intervals = append(intervals, []int{this.inner[0], this.inner[0]})
		for _, item := range this.inner {
			if n := len(intervals) - 1; intervals[n][1] >= item-1 {
				intervals[n][1] = item
			} else {
				intervals = append(intervals, []int{item, item})
			}
		}
	}
	return intervals
}

func main() {}
