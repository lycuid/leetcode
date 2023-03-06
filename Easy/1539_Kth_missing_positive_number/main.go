// https://leetcode.com/problems/kth-missing-positive-number/
package main

func findKthPositive(nums []int, k int) (unused int) {
	used := 0
	for i := 1; unused < k; i++ {
		if used < len(nums) && nums[used] == i {
			used++
		} else {
			unused++
		}
	}
	return used + unused
}

func main() {}
