// https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/
package main

import "sort"

func maxCount(banned []int, n int, maxSum int) (count int) {
	sort.Ints(banned)
	for i := 1; i <= n; i++ {
		if len(banned) > 0 && banned[0] == i {
			for len(banned) > 0 && banned[0] == i {
				banned = banned[1:]
			}
		} else {
			if maxSum-i < 0 {
				break
			}
			maxSum, count = maxSum-i, count+1
		}
	}
	return count
}

func main() {}
