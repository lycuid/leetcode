// https://leetcode.com/problems/reverse-words-in-a-string/
package main

func reverseWords(s string) (ret string) {
	for i, j := 0, 0; i < len(s); {
		for ; i < len(s) && s[i] == ' '; i++ {
		}
		for j = i; i < len(s) && s[i] != ' '; i++ {
		}
		if j < i {
			ret = " " + s[j:i] + ret
		}
	}
	return ret[1:]
}

func main() {}
