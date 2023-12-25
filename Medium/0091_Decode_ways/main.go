// https://leetcode.com/problems/decode-ways/
package main

var valid = map[string]bool{
	"1":  true,
	"2":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"11": true,
	"12": true,
	"13": true,
	"14": true,
	"15": true,
	"16": true,
	"17": true,
	"18": true,
	"19": true,
	"20": true,
	"21": true,
	"22": true,
	"23": true,
	"24": true,
	"25": true,
	"26": true,
}

func numDecodings(s string) int {
	cache, n := make([]int, len(s)+1), len(s)
	cache[0] = 1
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if sub := s[j:i]; valid[sub] {
				cache[i] += cache[j]
			}
		}
	}
	return cache[n]
}

func main() {}
