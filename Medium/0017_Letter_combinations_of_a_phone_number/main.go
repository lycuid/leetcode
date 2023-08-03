// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
package main

var letters = [10][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

func letterCombinations(digits string) (ret []string) {
	if len(digits) == 0 {
		return ret
	}
	if len(digits) == 1 {
		return letters[digits[0]-'0']
	}
	head, tail := digits[0]-'0', letterCombinations(digits[1:])
	for _, lst := range tail {
		for _, ch := range letters[head] {
			ret = append(ret, string(ch)+lst)
		}
	}
	return ret
}

func main() {}
