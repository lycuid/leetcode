// https://leetcode.com/problems/number-of-ways-to-split-array/
package main

func waysToSplitArray(nums []int) (count int) {
	max, min_nums := nums[0], make([]int, len(nums))
	min_nums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		min_nums[i] += (nums[i] + min_nums[i+1])
	}
	for i := 1; i < len(nums); i++ {
		if max >= min_nums[i] {
			count++
		}
		max += nums[i]
	}
	return count
}

func main() {}
