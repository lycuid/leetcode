// https://leetcode.com/problems/increasing-triplet-subsequence/
package main

func increasingTriplet(nums []int) bool {
	var (
		sub    [3]int
		cursor int
	)
	for i := 0; cursor < 3 && i < len(nums); i++ {
		if cursor == 0 || sub[cursor-1] < nums[i] {
			sub[cursor], cursor = nums[i], cursor+1
		} else {
			if sub[0] >= nums[i] {
				sub[0] = nums[i]
			} else if cursor >= 1 && sub[1] >= nums[i] {
				sub[1] = nums[i]
			}
		}
	}
	return cursor == 3
}

func main() {}
