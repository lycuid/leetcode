// https://leetcode.com/problems/subsets/
package main

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{nil}
	}
	subs := subsets(nums[1:])
	for i, n := 0, len(subs); i < n; i++ {
		subs = append(subs, append([]int{nums[0]}, subs[i]...))
	}
	return subs
}

func main() {}
