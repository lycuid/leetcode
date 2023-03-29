// https://leetcode.com/problems/reducing-dishes/
package main

import "sort"

func Aux(nums []int) (ret int) {
	for i := range nums {
		ret += nums[i] * (i + 1)
	}
	return ret
}

func maxSatisfaction(satisfaction []int) (ret int) {
	sort.Ints(satisfaction)
	for i, n := 0, len(satisfaction); i < n; i++ {
		if sum := Aux(satisfaction[i:]); ret < sum {
			ret = sum
		}
	}
	return ret
}

func main() {}
