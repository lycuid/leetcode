// https://leetcode.com/problems/different-ways-to-add-parentheses/
package main

import "strconv"

func diffWaysToCompute(expression string) (ret []int) {
	for i, op := range expression {
		if op == '+' || op == '-' || op == '*' {
			nums1 := diffWaysToCompute(expression[:i])
			nums2 := diffWaysToCompute(expression[i+1:])
			for _, n1 := range nums1 {
				for _, n2 := range nums2 {
					switch op {
					case '+':
						ret = append(ret, n1+n2)
					case '-':
						ret = append(ret, n1-n2)
					case '*':
						ret = append(ret, n1*n2)
					}
				}
			}
		}
	}
	if len(ret) == 0 {
		num, _ := strconv.Atoi(expression)
		ret = append(ret, num)
	}
	return ret
}

func main() {}
