// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k
package main

func minOperations(nums []int, k int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] += nums[i]
	}
	return nums[0] % k
}

func main() {}
