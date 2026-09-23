// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
package main

func minOperations(nums []int, x int) int {
	var (
		n     = len(nums)
		res   = n + 1
		cache = make([]int, n+1)
	)
	for i := range nums {
		cache[i+1] = cache[i] + nums[i]
	}
	if diff := cache[n] - x; diff == 0 {
		return n
	} else if diff > 0 {
		for i, j := 0, 0; i < n; i++ {
			for j <= i && cache[i+1]-cache[j] > diff {
				j++
			}
			if cache[i+1]-cache[j] == diff {
				res = min(res, n-(i-j+1))
			}
		}
	}
	if res > n {
		return -1
	}
	return res
}

func main() {}
