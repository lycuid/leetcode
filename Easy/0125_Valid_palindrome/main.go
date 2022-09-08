// https://leetcode.com/problems/valid-palindrome/
package main

func in(x, min, max int) bool {
	return x >= min && x <= max
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isPalindrome(s string) bool {
	var (
		s1 []byte
		l  int
	)

	for i := 0; i < len(s); i++ {
		if in(int(s[i]), 65, 90) || in(int(s[i]), 97, 122) || in(int(s[i]), 47, 57) {
			i1 := s[i]
			if in(int(i1), 97, 122) {
				i1 -= 32
			}
			s1 = append(s1, i1)
			l++
		}
	}

	for j := 0; j < l/2; j++ {
		if s1[j] != s1[l-j-1] {
			return false
		}
	}

	return true
}

func main() {}
