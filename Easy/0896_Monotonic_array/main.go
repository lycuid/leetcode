// https://leetcode.com/problems/monotonic-array/
package main

func isMonotonic(nums []int) bool {
	for i, j, s := 0, 1, 0; j < len(nums); i, j = i+1, j+1 {
		switch s {
		case -1:
			if nums[i] < nums[j] {
				return false
			}
		case 1:
			if nums[i] > nums[j] {
				return false
			}
		case 0:
			if nums[i] < nums[j] {
				s = 1
			} else if nums[i] > nums[j] {
				s = -1
			}
		}
	}
	return true
}

func main() {}
