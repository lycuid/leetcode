// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package main

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			left, lower, upper, right := 0, mid, mid, len(nums)-1
			if nums[left] == target {
				lower = left
			} else {
				for left+1 < lower {
					mid := (left + lower) / 2
					if nums[mid] == target {
						lower = mid
					} else {
						left = mid
					}
				}
			}
			if nums[right] == target {
				upper = right
			} else {
				for upper < right-1 {
					mid := (upper + right) / 2
					if nums[mid] == target {
						upper = mid
					} else {
						right = mid
					}
				}
			}
			return []int{lower, upper}
		}
	}
	return []int{-1, -1}
}

func main() {}
