// https://leetcode.com/problems/maximum-length-of-repeated-subarray/
package main

func findLength(nums1 []int, nums2 []int) (ret int) {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	cache := make([]int, len(nums2)+1)
	for i := range nums1 {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				cache[j+1] = 1 + cache[j]
			} else {
				cache[j+1] = 0
			}
			if cache[j+1] > ret {
				ret = cache[j+1]
			}
		}
	}
	return ret
}

func main() {}
