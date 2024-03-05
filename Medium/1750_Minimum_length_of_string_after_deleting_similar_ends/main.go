// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/
package main

func minimumLength(s string) int {
	i, j := 0, len(s)-1
	for i < j && s[i] == s[j] {
		tmp := s[i]
		for i < j && s[i] == tmp {
			i++
		}
		for i <= j && s[j] == tmp {
			j--
		}
	}
	return j - i + 1
}

func main() {}
