// https://leetcode.com/problems/largest-3-same-digit-number-in-string/
package main

func largestGoodInteger(num string) (digit string) {
	for i := 0; i <= len(num)-3; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if len(digit) == 0 || digit[0]-'0' < num[i]-'0' {
				digit = string(num[i])
			}
		}
	}
	return digit + digit + digit
}

func main() {}
