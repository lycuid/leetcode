// https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/
package main

func maxDistance(nums1 []int, nums2 []int) (res int) {
	for i, j := 0, 0; j < len(nums2); j = max(i, j+1) {
		if nums1[i] <= nums2[j] {
			res = max(res, j-i)
		}
		for i < len(nums1)-1 && nums1[i] > nums2[j] {
			i++
		}
	}
	return res
}

func main() {}
