// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
package main

func getNoZeroIntegers(n int) []int {
	ret := []int{0, n}
	contains0 := func(n int) bool {
		if n == 0 {
			return true
		}
		for ; n > 0; n /= 10 {
			if n%10 == 0 {
				return true
			}
		}
		return false
	}
	for ; ret[0] <= ret[1]; ret[0], ret[1] = ret[0]+1, ret[1]-1 {
		if !contains0(ret[0]) && !contains0(ret[1]) {
			break
		}
	}
	return ret
}

func main() {}
