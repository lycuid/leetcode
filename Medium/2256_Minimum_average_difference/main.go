// https://leetcode.com/problems/minimum-average-difference/
package main

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func minimumAverageDifference(nums []int) int {
	pre, n := make([]int, len(nums)+1), len(nums)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + nums[i-1]
	}
	post, min, ret := 0, pre[n]/n, n-1
	for i := n - 2; i >= 0; i-- {
		post += nums[i+1]
		if s := Abs((pre[i+1] / (i + 1)) - (post / (n - i - 1))); s <= min {
			min, ret = s, i
		}
	}
	return ret
}

func main() {}
