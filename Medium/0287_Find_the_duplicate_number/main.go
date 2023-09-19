// https://leetcode.com/problems/find-the-duplicate-number/
package main

func findDuplicate(nums []int) int {
	slow, fast, ptr := 0, 0, 0
	for {
		if slow, fast = nums[slow], nums[nums[fast]]; slow == fast {
			break
		}
	}
	for slow != ptr {
		slow, ptr = nums[slow], nums[ptr]
	}
	return slow
}

func main() {}
