// https://leetcode.com/problems/removing-stars-from-a-string/
package main

func removeStars(s string) string {
	bytes, l, r := []byte(s), 0, 0
	for ; r < len(bytes); l, r = l+1, r+1 {
		if bytes[r] == '*' {
			l -= 2
		} else {
			bytes[l] = bytes[r]
		}
	}
	return string(bytes[:l])
}

func main() {}
