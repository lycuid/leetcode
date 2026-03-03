// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/
package main

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	l := (1 << n) - 1
	if mid := (l >> 1) + 1; k > mid {
		return '0' + '1' - findKthBit(n-1, l-k+1)
	} else if k < mid {
		return findKthBit(n-1, k)
	}
	return '1'
}

func main() {}
