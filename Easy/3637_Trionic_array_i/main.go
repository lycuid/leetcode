// https://leetcode.com/problems/trionic-array-i/
package main

func accumulate(nums []int, fn func(int, int) bool) ([]int, bool) {
	if len(nums) == 0 {
		return nums, false
	}
	var dropped int
	for ; len(nums) > 1 && fn(nums[0], nums[1]); dropped++ {
		nums = nums[1:]
	}
	return nums, dropped > 0
}

func isTrionic(nums []int) bool {
	var (
		ok  bool
		inc = func(i, j int) bool { return i < j }
		dec = func(i, j int) bool { return i > j }
	)
	if nums, ok = accumulate(nums, inc); !ok {
		return false
	}
	if nums, ok = accumulate(nums, dec); !ok {
		return false
	}
	if nums, ok = accumulate(nums, inc); !ok {
		return false
	}
	return len(nums) == 1
}

func main() {}
