// https://leetcode.com/problems/combination-sum-ii/
package main

import "sort"

func Aux(nums []int, target int) (ret [][]int) {
	if target == 0 {
		return [][]int{nil}
	}
	if len(nums) == 0 || nums[0] > target {
		return nil
	}
	for i := 0; i < len(nums); {
		for _, val := range Aux(nums[i+1:], target-nums[i]) {
			ret = append(ret, append([]int{nums[i]}, val...))
		}
		for i++; i < len(nums) && nums[i-1] == nums[i]; i++ {
		}
	}
	return ret
}

func combinationSum2(candidates []int, target int) (ret [][]int) {
	sort.Ints(candidates)
	return Aux(candidates, target)
}

func main() {}
