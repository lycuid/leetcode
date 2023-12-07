// https://leetcode.com/problems/largest-odd-number-in-string/
package main

func largestOddNumber(num string) string {
	for len(num) > 0 && (num[len(num)-1]-'0')%2 == 0 {
		num = num[:len(num)-1]
	}
	return num
}

func main() {}
