// https://leetcode.com/problems/basic-calculator/
package main

func ConsumeNumber(s *string) (ret int) {
	for ; len(*s) > 0 && (*s)[0] >= '0' && (*s)[0] <= '9'; *s = (*s)[1:] {
		ret *= 10
		ret += int((*s)[0] - '0')
	}
	return ret
}

func Consume(s *string) (ret int) {
	var (
		Plus  = func(i, j int) int { return i + j }
		Minus = func(i, j int) int { return i - j }
	)
	for operation := Plus; len(*s) > 0; {
		switch (*s)[0] {
		case ' ':
			break
		case '-':
			operation = Minus
		case '+':
			operation = Plus
		case '(':
			*s = (*s)[1:]
			ret = operation(ret, Consume(s))
		case ')':
			goto EXIT
		default:
			ret = operation(ret, ConsumeNumber(s))
			continue
		}
		*s = (*s)[1:]
	}
EXIT:
	return ret
}

func calculate(s string) (ret int) { return Consume(&s) }

func main() {}
