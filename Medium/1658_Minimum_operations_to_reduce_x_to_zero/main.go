// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
package main

func minOperations(nums []int, x int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	sub, max, target := 0, -1, sum-x
	for i, j := 0, 0; j < len(nums); {
		for sub, j = sub+nums[j], j+1; i < j && sub > target; i++ {
			sub -= nums[i]
		}
		if n := j - i; n > max && sub == target {
			max = n
		}
	}
	if max == -1 {
		return -1
	}
	return len(nums) - max
}

func main() {}
