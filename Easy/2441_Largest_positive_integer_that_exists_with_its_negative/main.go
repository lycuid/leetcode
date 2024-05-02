// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/
package main

func findMaxK(nums []int) int {
	const N = 2001
	var cache [N]bool
	for _, num := range nums {
		cache[num+N/2] = true
	}
	for i := 0; i < N/2; i++ {
		if cache[i] && cache[N-1-i] {
			return N/2 - i
		}
	}
	return -1
}

func main() {}
