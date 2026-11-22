// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/
package main

func checkString(s string) bool {
	bindex := -1
	for i := range s {
		switch ch := s[i]; ch {
		case 'a':
			if bindex != -1 {
				return false
			}
		case 'b':
			bindex = i
		}
	}
	return true
}

func main() {}
