// https://leetcode.com/problems/combination-sum-ii/
package main

import "sort"

func Aux(nums []int, target int) (ret [][]int) {
	if target >= 0 {
		if target == 0 {
			return [][]int{nil}
		}
		for i := 0; i < len(nums); {
			for _, comb := range Aux(nums[i+1:], target-nums[i]) {
				ret = append(ret, append([]int{nums[i]}, comb...))
			}
			for i++; i < len(nums) && nums[i-1] == nums[i]; i++ {
			}
		}
	}
	return ret
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return Aux(candidates, target)
}

func main() {}
