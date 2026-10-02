// https://leetcode.com/problems/generate-parentheses/
package main

func generateParenthesis(n int) (res []string) {
	var solve func(string, int, int)
	solve = func(current string, open, close int) {
		if len(current) == n*2 {
			res = append(res, current)
			return
		}
		if open < n {
			solve(current+"(", open+1, close)
		}
		if close < open {
			solve(current+")", open, close+1)
		}
	}
	solve("", 0, 0)
	return res
}

func main() {}
