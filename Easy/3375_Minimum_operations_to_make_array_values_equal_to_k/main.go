// https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/
package main

func minOperations(nums []int, k int) int {
	for _, num := range nums {
		if num < k {
			return -1
		}
	}
	cache := make(map[int]bool)
	for _, num := range nums {
		if num != k {
			cache[num] = true
		}
	}
	return len(cache)
}

func main() {}
