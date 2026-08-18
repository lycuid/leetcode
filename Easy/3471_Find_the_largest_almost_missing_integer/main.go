// https://leetcode.com/problems/find-the-largest-almost-missing-integer/
package main

func largestInteger(nums []int, k int) int {
	var cache, index [51]int
	for i := range index {
		index[i] = -1
	}
	for i, num := range nums {
		left, right := i-index[num], len(nums)-i
		if window := left + right - 1; window >= k {
			cache[num] += min(left, right, k, window-k+1)
		}
		index[num] = i
	}
	for i := len(cache) - 1; i >= 0; i-- {
		if cache[i] == 1 {
			return i
		}
	}
	return -1
}

func main() {}
