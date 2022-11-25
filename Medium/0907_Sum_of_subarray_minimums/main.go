// https://leetcode.com/problems/sum-of-subarray-minimums/
package main

func sumSubarrayMins(nums []int) (ret int) {
	sum, stack := make([]int, len(nums)), []int{}
	for i := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			stack = stack[:len(stack)-1]
		}
		if n := len(stack); n > 0 {
			sum[i] = (i-stack[n-1])*nums[i] + sum[stack[n-1]]
		} else {
			sum[i] = (i + 1) * nums[i]
		}
		stack, sum[i] = append(stack, i), sum[i]%(1e9+7)
	}
	for i := range sum {
		ret = (ret + sum[i]) % (1e9 + 7)
	}
	return ret
}

func main() {}
