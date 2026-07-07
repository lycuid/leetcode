// https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/
package main

func sumAndMultiply(n int) int64 {
	var _n, x, s int64
	for ; n > 0; n /= 10 {
		if m := int64(n % 10); m > 0 {
			_n, s = _n*10+m, s+m
		}
	}
	for ; _n > 0; _n /= 10 {
		x = x*10 + (_n % 10)
	}
	return x * s
}

func main() {}
