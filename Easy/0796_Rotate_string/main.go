// https://leetcode.com/problems/rotate-string/
package main

func rotateString(s string, goal string) bool {
	if n := len(goal); len(s) == n {
		for s = s + s; len(s) > n; s = s[1:] {
			if s[:n] == goal {
				return true
			}
		}
	}
	return false
}

func main() {}
