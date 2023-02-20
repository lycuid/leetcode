// https://leetcode.com/problems/search-insert-position/
package main

func searchInsert(nums []int, target int) (l int) {
	for r := len(nums) - 1; l < r; {
		if mid := (l + r) / 2; nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			l, r = mid, mid
		}
	}
	if nums[l] < target {
		l++
	}
	return l
}

func main() {}
