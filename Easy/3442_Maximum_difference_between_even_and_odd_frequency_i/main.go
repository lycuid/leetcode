// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/
package main

func maxDifference(s string) int {
	cache := [26]int{}
	for i := range s {
		cache[s[i]-'a']++
	}
	a := [2]int{len(s), 0}
	for _, freq := range cache {
		if freq != 0 {
			if freq%2 == 0 {
				a[0] = min(a[0], freq)
			} else {
				a[1] = max(a[1], freq)
			}
		}
	}
	return a[1] - a[0]
}

func main() {}
