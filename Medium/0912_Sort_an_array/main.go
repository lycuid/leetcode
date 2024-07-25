// https://leetcode.com/problems/sort-an-array/
package main

func Merge(nums []int, mid int) {
	nums1, nums2, cache, i := nums[:mid], nums[mid:], make([]int, len(nums)), 0
	for ; len(nums1) > 0 && len(nums2) > 0; i++ {
		if nums1[0] < nums2[0] {
			cache[i], nums1 = nums1[0], nums1[1:]
		} else {
			cache[i], nums2 = nums2[0], nums2[1:]
		}
	}
	for ; len(nums1) > 0; i++ {
		cache[i], nums1 = nums1[0], nums1[1:]
	}
	for ; len(nums2) > 0; i++ {
		cache[i], nums2 = nums2[0], nums2[1:]
	}
	for i, num := range cache {
		nums[i] = num
	}
}

func Sort(nums []int) []int {
	if n := len(nums); n > 1 {
		mid := n / 2
		Sort(nums[:mid])
		Sort(nums[mid:])
		Merge(nums, mid)
	}
	return nums
}

func sortArray(nums []int) []int {
	return Sort(nums)
}

func main() {}
