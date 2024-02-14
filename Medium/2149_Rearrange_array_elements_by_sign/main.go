// https://leetcode.com/problems/rearrange-array-elements-by-sign/
package main

func rearrangeArray(nums []int) []int {
	ret := make([]int, len(nums))
	for i, p, n := 0, 0, 1; i < len(nums); i++ {
		if nums[i] >= 0 {
			ret[p], p = nums[i], p+2
		} else {
			ret[n], n = nums[i], n+2
		}
	}
	return ret
}

func main() {}
