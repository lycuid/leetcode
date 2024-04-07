// https://leetcode.com/problems/valid-parenthesis-string/
package main

func checkValidString(s string) bool {
	var left, misc []int
	for i, ch := range s {
		switch ch {
		case '(':
			left = append(left, i)
			break
		case ')':
			if n := len(left); n > 0 {
				left = left[:n-1]
			} else if m := len(misc); m > 0 {
				misc = misc[:m-1]
			} else {
				return false
			}
			break
		default:
			misc = append(misc, i)
			break
		}
	}
	for len(left) > 0 && len(misc) > 0 {
		for len(misc) > 0 && misc[0] <= left[0] {
			misc = misc[1:]
		}
		if len(misc) > 0 {
			left, misc = left[1:], misc[1:]
		}
	}
	return len(left) == 0
}

func main() {}
