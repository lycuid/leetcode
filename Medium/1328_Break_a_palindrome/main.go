// https://leetcode.com/problems/break-a-palindrome/
package main

func breakPalindrome(palindrome string) string {
	if n := len(palindrome); n > 1 {
		bs, i := []byte(palindrome), 0
		for ; i < n && bs[i] == 'a'; i++ {
		}
		if i < n && (n%2 == 0 || i != n/2) {
			bs[i] = 'a'
		} else {
			bs[n-1] = 'b'
		}
		return string(bs)
	}
	return ""
}

func main() {}
