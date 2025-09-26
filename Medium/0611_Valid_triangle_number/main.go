// https://leetcode.com/problems/valid-triangle-number/
package main

import "sort"

func triangleNumber(nums []int) (count int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			target, l, r := nums[i]+nums[j], j+1, len(nums)
			for l < r {
				if mid := (l + r) / 2; nums[mid] < target {
					l = mid + 1
				} else {
					r = mid
				}
			}
			count += l - j - 1
		}
	}
	return count
}

func main() {}
