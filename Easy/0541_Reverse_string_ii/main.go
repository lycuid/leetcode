// https://leetcode.com/problems/reverse-string-ii/
package main

func reverseStr(s string, k int) string {
	chars := []byte(s)
	for i, n := 0, len(chars); i < n; i += 2 * k {
		l, r := i, i+k-1
		if r >= n {
			r = n - 1
		}
		for ; l < r; l, r = l+1, r-1 {
			chars[l], chars[r] = chars[r], chars[l]
		}
	}
	return string(chars)
}

func main() {}
