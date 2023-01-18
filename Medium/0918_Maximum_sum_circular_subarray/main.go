// https://leetcode.com/problems/maximum-sum-circular-subarray/
package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	var (
		curr, sum [2]int
		total     int
	)
	const Max, Min = 0, 1
	sum[Max], sum[Min] = math.MinInt, math.MaxInt
	for _, num := range nums {
		if curr[Max] += num; curr[Max] < num {
			curr[Max] = num
		}
		if sum[Max] < curr[Max] {
			sum[Max] = curr[Max]
		}
		if curr[Min] += num; curr[Min] > num {
			curr[Min] = num
		}
		if sum[Min] > curr[Min] {
			sum[Min] = curr[Min]
		}
		total += num
	}
	if total == sum[Min] {
		return sum[Max]
	}
	if sum[Max] < total-sum[Min] {
		sum[Max] = total - sum[Min]
	}
	return sum[Max]
}

func main() {}
