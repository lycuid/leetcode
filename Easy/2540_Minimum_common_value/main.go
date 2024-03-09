// https://leetcode.com/problems/minimum-common-value/
package main

func getCommon(nums1 []int, nums2 []int) int {
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			return nums1[i]
		}
	}
	return -1
}

func main() {}
