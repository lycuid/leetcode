// https://leetcode.com/problems/smallest-number-with-all-set-bits/
package main

func smallestNumber(n int) (num int) {
	for num = 1; num < n; {
		num = num<<1 | 1
	}
	return num
}

func main() {}
