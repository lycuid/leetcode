// https://leetcode.com/problems/first-missing-positive/
package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] <= 0 || nums[i] > n || nums[i] == nums[nums[i]-1] {
			if nums[i] != i+1 {
				nums[i] = -1
			}
			i++
			continue
		}
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}
	for i, val := range nums {
		if val == -1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {}
