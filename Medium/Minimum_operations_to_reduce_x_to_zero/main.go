// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
package main

func minOperations(nums []int, x int) (res int) {
	n, target, sum := len(nums), -x, 0
	res = n + 1
	for i := 0; i < n; i++ {
		target += nums[i]
	}
	for i, j := 0, 0; i < n; i++ {
		sum += nums[i]
		for ; sum > target && j <= i; j++ {
			sum -= nums[j]
		}
		if sum == target && res > n-(i-j+1) {
			res = n - (i - j + 1)
		}
	}
	if res > n {
		return -1
	}
	return res
}

func main() {}
