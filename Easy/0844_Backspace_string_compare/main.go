// https://leetcode.com/problems/backspace-string-compare/
package main

func backspaceCompare(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	i, j := 0, 0
	for k := i; ; i, k = i+1, k+1 {
		for k < len(s1) && s1[k] == '#' {
			if i, k = i-1, k+1; i < 0 {
				i = 0
			}
		}
		if k >= len(s1) {
			break
		}
		s1[i] = s1[k]
	}
	for k := j; ; j, k = j+1, k+1 {
		for k < len(s2) && s2[k] == '#' {
			if j, k = j-1, k+1; j < 0 {
				j = 0
			}
		}
		if k >= len(s2) {
			break
		}
		s2[j] = s2[k]
	}
	return string(s1[:i]) == string(s2[:j])
}

func main() {}
