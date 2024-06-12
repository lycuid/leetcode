// https://leetcode.com/problems/sort-colors/
package main

func Merge(nums []int) {
	n := len(nums)
	nums1, nums2, tmp := nums[:n/2], nums[n/2:], make([]int, 0, n)
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] < nums2[0] {
			tmp, nums1 = append(tmp, nums1[0]), nums1[1:]
		} else {
			tmp, nums2 = append(tmp, nums2[0]), nums2[1:]
		}
	}
	for ; len(nums1) > 0; nums1 = nums1[1:] {
		tmp = append(tmp, nums1[0])
	}
	for ; len(nums2) > 0; nums2 = nums2[1:] {
		tmp = append(tmp, nums2[0])
	}
	for i := 0; i < n; i++ {
		nums[i] = tmp[i]
	}
}

func sortColors(nums []int) {
	if n := len(nums); n > 1 {
		mid := n / 2
		sortColors(nums[:mid])
		sortColors(nums[mid:])
		Merge(nums)
	}
}

func main() {}
