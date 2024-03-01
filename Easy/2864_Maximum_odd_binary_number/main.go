// https://leetcode.com/problems/maximum-odd-binary-number/
package main

func maximumOddBinaryNumber(s string) string {
	str, j := []byte(s), 0
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			str[i], str[j], j = str[j], str[i], j+1
		}
	}
	if j > 0 {
		str[j-1], str[len(str)-1] = str[len(str)-1], str[j-1]
	}
	return string(str)
}

func main() {}
