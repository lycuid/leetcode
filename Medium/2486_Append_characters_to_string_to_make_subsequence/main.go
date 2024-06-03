// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/
package main

func appendCharacters(s, t string) int {
	for ; len(s) > 0 && len(t) > 0; s = s[1:] {
		if s[0] == t[0] {
			t = t[1:]
		}
	}
	return len(t)
}

func main() {}
