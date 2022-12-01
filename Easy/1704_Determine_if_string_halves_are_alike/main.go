// https://leetcode.com/problems/determine-if-string-halves-are-alike/
package main

func halvesAreAlike(s string) bool {
	var n int
	for i, mid := 0, len(s)/2; i < mid*2; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < mid {
				n++
			} else {
				n--
			}
		default:
			break
		}
	}
	return n == 0
}

func main() {}
