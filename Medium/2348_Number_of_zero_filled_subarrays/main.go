// https://leetcode.com/problems/number-of-zero-filled-subarrays/
package main

func zeroFilledSubarray(nums []int) (ret int64) {
	for i, j := 0, 0; i < len(nums); i = j {
		for ; i < len(nums) && nums[i] != 0; i++ {
		}
		for j = i; j < len(nums) && nums[j] == 0; j++ {
		}
		ret += int64(((j - i) * ((j - i) + 1)) / 2)
	}
	return ret
}

func main() {}
