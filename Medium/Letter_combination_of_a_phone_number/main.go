// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

package main

var letterMappings = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func head(str string) (s byte) {
	if len(str) > 0 {
		return str[0]
	}
	return s
}

func tail(str string) (s string) {
	if len(str) > 1 {
		return str[1:]
	}
	return s
}

func letterCombinations(digits string) (s []string) {
	h := head(digits)
	t := tail(digits)

	if h == '0' {
		return s
	}

	if len(t) == 0 {
		for _, letter := range letterMappings[h] {
			s = append(s, string(letter))
		}
		return s
	}

	for _, letter := range letterMappings[h] {
		for _, tailString := range letterCombinations(t) {
			s = append(s, string(letter)+tailString)
		}
	}

	return s
}

func main() {}
