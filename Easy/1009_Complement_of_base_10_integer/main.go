// https://leetcode.com/problems/complement-of-base-10-integer/
package main

func bitwiseComplement(n int) int {
	mask := 2
	for mask <= n {
		mask <<= 1
	}
	return n ^ (mask - 1)
}

func main() {}
