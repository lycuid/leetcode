// https://leetcode.com/problems/binary-search/
package main

func search(nums []int, target int) (l int) {
	for r := len(nums) - 1; l < r; {
		if mid := (l + r) / 2; nums[mid] >= target {
			r = mid
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	if nums[l] != target {
		l = -1
	}
	return l
}

func main() {}
