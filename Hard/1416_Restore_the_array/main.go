// https://leetcode.com/problems/restore-the-array/
package main

const MOD = 1e9 + 7

func numberOfArrays(s string, k int) int {
	cache, n, digits := make([]int, len(s)+1), len(s), 0
	cache[0] = 1
	for i := k; i > 0; i /= 10 {
		digits++
	}
	for i, num := 1, 0; i <= n; i, num = i+1, 0 {
		for j, mul := i-1, 1; j >= 0 && j >= i-digits; j, mul = j-1, mul*10 {
			if ch := int(s[j] - '0'); ch != 0 {
				if num += mul * ch; num > k {
					break
				}
				cache[i] = (cache[i] + cache[j]) % MOD
			}
		}
	}
	return cache[n]
}

func main() {}
