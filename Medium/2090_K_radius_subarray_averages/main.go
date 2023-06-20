// https://leetcode.com/problems/k-radius-subarray-averages/
package main

func getAverages(nums []int, k int) []int {
	ret := make([]int, len(nums))
	for i := range ret {
		ret[i] = -1
	}
	sum, d, n := 0, 2*k+1, len(nums)
	for i := 0; i < n && i < d-1; i++ {
		sum += nums[i]
	}
	for i := k; i+k < n; i++ {
		if sum += nums[i+k]; i-k-1 >= 0 {
			sum -= nums[i-k-1]
		}
		ret[i] = sum / d
	}
	return ret
}

func main() {}
