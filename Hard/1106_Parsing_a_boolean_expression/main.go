// https://leetcode.com/problems/parsing-a-boolean-expression/
package main

func parseExpr(expr string) (bool, string) {
	trimLeft := func(s string) string {
		for len(s) > 0 && s[0] == ' ' {
			s = s[1:]
		}
		return s
	}

	parse := func(e string, fn func(bool, bool) bool) (bool, string) {
		var acc, val bool
		acc, e = parseExpr(trimLeft(e[1:]))
		for e = trimLeft(e); len(e) > 0 && e[0] == ','; e = trimLeft(e) {
			val, e = parseExpr(e[1:])
			acc = fn(acc, val)
		}
		return acc, e[1:]
	}

	switch expr[0] {
	case '&':
		return parse(expr[1:], func(i, j bool) bool { return i && j })
	case '|':
		return parse(expr[1:], func(i, j bool) bool { return i || j })
	case '!':
		tmp, e := parse(expr[1:], func(i, j bool) bool { return i })
		return !tmp, e
	case 't':
		return true, expr[1:]
	case 'f':
		return false, expr[1:]
	}
	return false, expr
}

func parseBoolExpr(expression string) bool {
	ret, _ := parseExpr(expression)
	return ret
}

func main() {}
