// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
package main

func longestSubarray(nums []int, limit int) (ret int) {
	var (
		smaller = make([]int, 0, len(nums))
		bigger  = make([]int, 0, len(nums))
		l       = 0
	)
	for r, num := range nums {
		for n := len(smaller) - 1; n >= 0 && num < smaller[n]; n = len(smaller) - 1 {
			smaller = smaller[:n]
		}
		smaller = append(smaller, num)
		for n := len(bigger) - 1; n >= 0 && num > bigger[n]; n = len(bigger) - 1 {
			bigger = bigger[:n]
		}
		bigger = append(bigger, num)
		for ; bigger[0]-smaller[0] > limit; l++ {
			if bigger[0] == nums[l] {
				bigger = bigger[1:]
			}
			if smaller[0] == nums[l] {
				smaller = smaller[1:]
			}
		}
		if diff := r - l + 1; diff > ret {
			ret = diff
		}
	}
	return ret
}

func main() {}
