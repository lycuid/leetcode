// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
package main

func concatenatedBinary(n int) (val int) {
	const MOD = 1e9 + 7
	for i, shift := 1, 0; i <= n; i++ {
		for 1<<shift <= i {
			shift++
		}
		val = ((val << shift) | i) % MOD
	}
	return val
}

func main() {}
