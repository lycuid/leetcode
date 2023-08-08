// https://leetcode.com/problems/search-in-rotated-sorted-array/
package main

func search(nums []int, target int) int {
	n := len(nums)
	var l, r int
	for l1 := n - 1; l < l1; {
		mid := (l + l1) / 2
		if nums[mid] >= nums[l] && nums[mid] > nums[l1] {
			l = mid + 1
		} else if nums[mid] < nums[l1] {
			l1 = mid
		} else {
			l, l1 = mid, mid
		}
	}
	for r = l - 1 + n; l < r; {
		if mid := ((l + r) / 2); nums[mid%n] > target {
			r = mid
		} else if nums[mid%n] < target {
			l = mid + 1
		} else {
			l, r = mid, mid
		}
	}
	if l %= n; nums[l] != target {
		l = -1
	}
	return l
}

func main() {}
