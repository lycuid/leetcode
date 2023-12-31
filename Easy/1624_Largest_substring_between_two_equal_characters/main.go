// https://leetcode.com/problems/largest-substring-between-two-equal-characters/
package main

func maxLengthBetweenEqualCharacters(s string) int {
	var (
		max   = -1
		cache [26]int
	)
	for i := range s {
		if n := s[i] - 'a'; cache[n] == 0 {
			cache[n] = i + 1
		} else if m := i - cache[n]; m > max {
			max = m
		}
	}
	return max
}

func main() {}
