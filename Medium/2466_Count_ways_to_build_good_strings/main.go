// https://leetcode.com/problems/cache-ways-to-build-good-strings/
package main

const MOD = 1e9 + 7

func countGoodStrings(low int, high int, zero int, one int) (ret int) {
	cache := make([]int, high+1)
	cache[0] = 1
	for i := 1; i <= high; i++ {
		if i-zero >= 0 {
			cache[i] = (cache[i] + cache[i-zero]) % MOD
		}
		if i-one >= 0 {
			cache[i] = (cache[i] + cache[i-one]) % MOD
		}
		if i >= low {
			ret = (ret + cache[i]) % MOD
		}
	}
	return ret
}

func main() {}
