// https://leetcode.com/problems/get-equal-substrings-within-budget/
package main

func equalSubstring(s string, t string, maxCost int) (ret int) {
	cache := make([]int, len(s))
	for i := range cache {
		if cache[i] = int(s[i]) - int(t[i]); cache[i] < 0 {
			cache[i] = -cache[i]
		}
	}
	for i, j, cost := 0, 0, 0; i < len(cache); i++ {
		for ; j < len(cache) && cost+cache[j] <= maxCost; j++ {
			cost += cache[j]
		}
		if sub := j - i; ret < sub {
			ret = sub
		}
		cost -= cache[i]
	}
	return ret
}

func main() {}
