// https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
package main

func canMakeSubsequence(str1 string, str2 string) bool {
	for len(str1) > 0 && len(str2) > 0 {
		ch := str2[0] - 'a'
		for len(str1) > 0 && str1[0]-'a' != ch && str1[0]-'a' != (ch+25)%26 {
			str1 = str1[1:]
		}
		if len(str1) > 0 {
			str1, str2 = str1[1:], str2[1:]
		}
	}
	return len(str2) == 0
}

func main() {}
