// https://leetcode.com/problems/power-of-four/
package main

func isPowerOfFour(n int) bool {
	if n&(n-1) > 0 {
		return false
	}
	for n >= 4 {
		n >>= 2
	}
	return n == 1
}

func main() {}
