// https://leetcode.com/problems/binary-subarrays-with-sum/
package main

func Aux(nums []int, goal int) (count int) {
	if goal >= 0 {
		for i, j, sum := 0, 0, 0; j < len(nums); j++ {
			for sum += nums[j]; sum > goal; i++ {
				sum -= nums[i]
			}
			count += j - i + 1
		}
	}
	return count
}

func numSubarraysWithSum(nums []int, goal int) int {
	return Aux(nums, goal) - Aux(nums, goal-1)
}

func main() {}
