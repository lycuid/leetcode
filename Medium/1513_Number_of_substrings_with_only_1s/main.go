// https://leetcode.com/problems/number-of-substrings-with-only-1s/
package main

func numSub(s string) (count int) {
	for i, j := 0, 0; i < len(s); i++ {
		if s[i] == '0' {
			j = i + 1
		}
		count = (count + i - j + 1) % (1e9 + 7)
	}
	return count
}

func main() {}
