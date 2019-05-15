// https://leetcode.com/problems/valid-parentheses/

package main

// Stack stack implementation
type Stack struct {
	stack []rune
}

func (s Stack) get() rune {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) add(i rune) {
	s.stack = append(s.stack, i)
}

func (s *Stack) pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s Stack) length() int {
	return len(s.stack)
}

func isValid(s string) bool {
	openers := map[rune]bool{40: true, 91: true, 123: true}
	closingMaps := map[rune]rune{
		41:  40,
		93:  91,
		125: 123,
	}

	stack := Stack{[]rune{}}

	for _, i := range s {
		if _, ok := openers[i]; ok {
			stack.add(i)
			continue
		}

		if stack.length() > 0 && stack.get() == closingMaps[i] {
			stack.pop()
			continue
		}

		return false
	}

	return stack.length() == 0
}

func main() {}
