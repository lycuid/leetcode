// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
package main

func findMin(nums []int) (l int) {
	for r := len(nums) - 1; l < r; {
		if mid := l + (r-l)>>1; nums[mid] <= nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}

func main() {}
