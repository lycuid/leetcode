// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/
package main

func findUnsortedSubarray(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end && (nums[start] <= nums[start+1] || nums[end] >= nums[end-1]) {
		if nums[start] <= nums[start+1] {
			start++
		}
		if nums[end] >= nums[end-1] {
			end--
		}
	}
	if start >= end {
		return 0
	}
	low, high := nums[end], nums[start]
	for i := start; i < end; i++ {
		if nums[i] < low {
			low = nums[i]
		}
		if nums[i] > high {
			high = nums[i]
		}
	}
	for start > 0 && nums[start-1] > low {
		start--
	}
	for end < len(nums)-1 && high > nums[end+1] {
		end++
	}
	return end - start + 1
}

func main() {}
