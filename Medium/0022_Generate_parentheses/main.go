// https://leetcode.com/problems/generate-parentheses/

package main

func solve(result *[]string, current string, opens, closes, n int) {
	if len(current) == n*2 {
		*result = append(*result, current)
		return
	}

	if opens < n {
		solve(result, current+"(", opens+1, closes, n)
	}

	if closes < opens {
		solve(result, current+")", opens, closes+1, n)
	}
}

func generateParenthesis(n int) (parentheses []string) {
	if n == 0 {
		return []string{""}
	}

	solve(&parentheses, "", 0, 0, n)
	return parentheses
}

/*
* Perfectly good solution but doesnt work,
* as the test cases are literally dumb.
 */

// func generateParenthesis(n int) (parentheses []string) {
// 	if n < 1 {
// 		return []string{""}
// 	}
// 	if n == 1 {
// 		return []string{"()"}
// 	}

// 	for _, paren := range generateParenthesis(n - 1) {
// 		parentheses = append(parentheses, "("+paren+")")

// 		left := "()" + paren
// 		right := paren + "()"

// 		parentheses = append(parentheses, right)
// 		if left != right {
// 			parentheses = append(parentheses, left)
// 		}
// 	}

// 	return parentheses
// }

func main() {}
