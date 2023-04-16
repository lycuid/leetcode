// https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
package main

const MOD = 1e9 + 7

func numWays(words []string, target string) int {
	freq, n := make([][26]int, len(words[0])), len(target)
	for _, word := range words {
		for i, ch := range word {
			freq[i][ch-'a']++
		}
	}
	cache := make([]int, n+1)
	cache[0] = 1
	for i := range words[0] {
		for j := n - 1; j >= 0; j-- {
			cache[j+1] += cache[j] * freq[i][target[j]-'a'] % MOD
			cache[j+1] %= MOD
		}
	}
	return cache[n]
}

func main() {}
