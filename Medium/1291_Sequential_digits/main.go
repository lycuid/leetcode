// https://leetcode.com/problems/sequential-digits/
package main

func from_digits(digits, start int) (n int, err bool) {
	if 9-start+1 < digits {
		return n, true
	}
	for i := start; digits > 0; i, digits = i+1, digits-1 {
		n = n*10 + i
	}
	return n, false
}

func digits(n int) (count int) {
	for ; n > 0; count++ {
		n /= 10
	}
	return count
}

func sequentialDigits(low int, high int) (ret []int) {
	for ld, hd := digits(low), digits(high); ld <= hd; ld++ {
		for i := 1; i <= 9; i++ {
			if n, err := from_digits(ld, i); err || n > high {
				break
			} else if n >= low {
				ret = append(ret, n)
			}
		}
	}
	return ret
}

func main() {}
