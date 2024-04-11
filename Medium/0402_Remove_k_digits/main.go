// https://leetcode.com/problems/remove-k-digits/
package main

func removeKdigits(num string, k int) string {
	stack := []byte{'0'}
	for i := range num {
		for ; k > 0 && stack[len(stack)-1] > num[i]; k-- {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:]
	}
	return string(stack)
}

func main() {}
