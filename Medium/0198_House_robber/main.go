// https://leetcode.com/problems/house-robber/
package main

func rob(nums []int) int {
	for i, fst, snd := len(nums)-1, 0, 0; i >= 0; i-- {
		if nums[i] += snd; nums[i] < fst {
			nums[i] = fst
		}
		fst, snd = nums[i], fst
	}
	return nums[0]
}

func main() {}
