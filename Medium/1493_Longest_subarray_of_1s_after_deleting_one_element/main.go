// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
package main

func longestSubarray(nums []int) (ret int) {
	for l, r, zeroes := 0, 0, 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zeroes++
		}
		for ; zeroes > 1; l++ {
			if nums[l] == 0 {
				zeroes--
			}
		}
		if n := r - l + 1 - zeroes; n > ret {
			ret = n
		}
	}
	if ret == len(nums) {
		return ret - 1
	}
	return ret
}

func main() {}
