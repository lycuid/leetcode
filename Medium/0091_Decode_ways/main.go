// https://leetcode.com/problems/decode-ways/
package main

func numDecodings(s string) (ret int) {
	next1, next2 := 1, 1
	if s[len(s)-1]-'0' != 0 {
		ret = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		ret, next1, next2 = 0, ret, next1
		if s[i]-'0' != 0 {
			ret += next1
		}
		if s[i]-'0' == 1 || (s[i]-'0' == 2 && s[i+1]-'0' <= 6) {
			ret += next2
		}
	}
	return ret
}

func main() {}
