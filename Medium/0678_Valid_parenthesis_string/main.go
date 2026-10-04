// https://leetcode.com/problems/valid-parenthesis-string/
package main

func checkValidString(s string) bool {
	open, misc := make([]int, 0, len(s)), make([]int, 0, len(s))
	for i, ch := range s {
		switch ch {
		case '(':
			open = append(open, i)
		case '*':
			misc = append(misc, i)
		case ')':
			if n := len(open); n > 0 {
				open = open[:n-1]
			} else if m := len(misc); m > 0 {
				misc = misc[:m-1]
			} else {
				return false
			}
		}
	}
	for ; len(open) > 0; open = open[1:] {
		for len(misc) > 0 && misc[0] < open[0] {
			misc = misc[1:]
		}
		if len(misc) == 0 {
			break
		}
		misc = misc[1:]
	}
	return len(open) == 0
}

func main() {}
