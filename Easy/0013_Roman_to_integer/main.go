// https://leetcode.com/problems/roman-to-integer/
package main

func romanToInt(s string) (ret int) {
	romans := []int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i, cur, prev := len(s)-1, 0, 0; i >= 0; i, prev = i-1, cur {
		cur = romans[s[i]]
		if ret = ret + cur; cur < prev {
			ret -= (2 * cur)
		}
	}
	return ret
}

func main() {}
