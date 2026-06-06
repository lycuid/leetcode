// https://leetcode.com/problems/left-and-right-sum-differences/
package main

func leftRightDifference(nums []int) []int {
	var left, right int
	for _, num := range nums {
		right += num
	}
	for i, num := range nums {
		right -= num
		if nums[i] = left - right; nums[i] < 0 {
			nums[i] = -nums[i]
		}
		left += num
	}
	return nums
}

func main() {}
