// https://leetcode.com/problems/summary-ranges/
package main

import "fmt"

func summaryRanges(nums []int) (ret []string) {
	if len(nums) > 0 {
		interval := [2]int{nums[0], nums[0]}
		AddToInterval := func() {
			if interval[0] == interval[1] {
				ret = append(ret, fmt.Sprintf("%d", interval[0]))
			} else {
				ret = append(ret, fmt.Sprintf("%d->%d", interval[0], interval[1]))
			}
		}
		for i := 1; i < len(nums); i++ {
			if nums[i] == interval[1]+1 {
				interval[1] = nums[i]
			} else {
				AddToInterval()
				interval[0], interval[1] = nums[i], nums[i]
			}
		}
		AddToInterval()
	}
	return ret
}

func main() {}
