// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
package main

import "fmt"

var chars = [10][]string{
	{}, {}, {"a", "b", "c"}, {"d", "e", "f"},
	{"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"},
	{"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}

func Combine(xs, ys *[]string) (ret []string) {
	if len(*ys) == 0 {
		return *xs
	}
	for _, x := range *xs {
		for _, y := range *ys {
			ret = append(ret, fmt.Sprintf("%s%s", x, y))
		}
	}
	return ret
}

func letterCombinations(digits string) (ret []string) {
	for i := len(digits) - 1; i >= 0; i-- {
		ret = Combine(&chars[digits[i]-'0'], &ret)
	}
	return ret
}

func main() {}
