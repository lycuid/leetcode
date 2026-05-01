// https://leetcode.com/problems/rotate-function/
package main

func maxRotateFunction(nums []int) int {
	var (
		n      = len(nums)
		cache  = make([]int, n)
		prefix = make([]int, 1, n+1)
	)
	for i, num := range nums {
		prefix = append(prefix, prefix[i]+num)
	}
	for i, num := range nums {
		cache[0] += i * num
	}
	for i := 1; i < n; i++ {
		cache[i] = cache[i-1] - (prefix[n] - prefix[i] + prefix[i-1]) + nums[i-1]*(n-1)
	}
	res := cache[0]
	for _, num := range cache {
		res = max(res, num)
	}
	return res
}

func main() {}
