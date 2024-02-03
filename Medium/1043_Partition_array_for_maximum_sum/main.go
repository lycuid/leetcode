// https://leetcode.com/problems/partition-array-for-maximum-sum/
package main

func Max(num int, nums ...int) int {
	for i := range nums {
		if nums[i] > num {
			num = nums[i]
		}
	}
	return num
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	cache := make([]int, n+1)
	for i := range arr {
		for j := 0; j < k && i-j >= 0; j++ {
			cache[i+1] = Max(cache[i+1], Max(arr[i], arr[i-j:i]...)*(j+1)+cache[i-j])
		}
	}
	return cache[n]
}

func main() {}
