// https://leetcode.com/problems/repeated-substring-pattern/
package main

func Aux(sub, str string) bool {
	for i, n := 0, len(sub); i+n <= len(str); i += n {
		if str[i:i+n] != sub {
			return false
		}
	}
	return true
}

func repeatedSubstringPattern(s string) bool {
	for i, n := 1, len(s); i <= n/2; i++ {
		if n%i == 0 && Aux(s[:i], s[i:]) {
			return true
		}
	}
	return false
}

func main() {}
