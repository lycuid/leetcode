// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k/
package main

import "math/bits"

func minOperations(nums []int, k int) int {
	for _, num := range nums {
		k ^= num
	}
	return bits.OnesCount(uint(k))
}

func main() {}
