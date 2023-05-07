// https://leetcode.com/problems/basic-calculator/
package main

func ParseNumber(s *string) (num int) {
	for ; len(*s) > 0 && (*s)[0] >= '0' && (*s)[0] <= '9'; *s = (*s)[1:] {
		num = num*10 + int((*s)[0]-'0')
	}
	return num
}

func Eval(s *string) (acc int) {
	Plus, Minus := func(i int) { acc += i }, func(i int) { acc -= i }
	for Action := Plus; len(*s) > 0; {
		switch (*s)[0] {
		case '+':
			Action = Plus
		case '-':
			Action = Minus
		case '(':
			*s = (*s)[1:]
			Action(Eval(s))
		case ')':
			goto DONE
		case ' ':
		default:
			Action(ParseNumber(s))
			continue
		}
		*s = (*s)[1:]
	}
DONE:
	return acc
}

func calculate(s string) int { return Eval(&s) }

func main() {}
