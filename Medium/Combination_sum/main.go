// https://leetcode.com/problems/combination-sum/
package main

func dfs(start, target int, nums []int) (ret [][]int) {
	for i := start; i < len(nums); i++ {
		if nums[i] == target {
			ret = append(ret, []int{nums[i]})
		}
		if nums[i] < target {
			for _, val := range dfs(i, target-nums[i], nums) {
				val = append([]int{nums[i]}, val...)
				ret = append(ret, val)
			}
		}
	}
	return ret
}

func combinationSum(nums []int, target int) [][]int {
	return dfs(0, target, nums)
}

func main() {}
