// https://leetcode.com/problems/wiggle-subsequence/
package main

func wiggleMaxLength(nums []int) int {
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			up = down + 1
		}
		if nums[i-1] > nums[i] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}

func main() {}
