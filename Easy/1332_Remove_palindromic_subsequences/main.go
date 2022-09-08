// https://leetcode.com/problems/remove-palindromic-subsequences/
package main

func removePalindromeSub(str string) (ret int) {
	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		if str[start] != str[end] {
			return 2
		}
	}
	return 1
}

func main() {}
