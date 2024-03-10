// https://leetcode.com/problems/intersection-of-two-arrays/
package main

import "sort"

func intersection(nums1 []int, nums2 []int) (ret []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			if n := len(ret); n == 0 || ret[n-1] != nums1[i] {
				ret = append(ret, nums1[i])
			}
			i, j = i+1, j+1
		}
	}
	return ret
}

func main() {}
