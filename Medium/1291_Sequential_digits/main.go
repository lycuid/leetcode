// https://leetcode.com/problems/sequential-digits/
package main

func solver(high int) func(int) int {
	digits := func(n int) (res int) {
		for ; n > 0; n /= 10 {
			res++
		}
		return res
	}
	return func(low int) int {
		for l, r := digits(low), digits(high); l <= r; l++ {
			for i := 1; i <= 10-l; i++ {
				var res int
				for j := 0; j < l; j++ {
					res = res*10 + i + j
				}
				if res > high {
					return 0
				}
				if res >= low {
					return res
				}
			}
		}
		return 0
	}
}

func sequentialDigits(low int, high int) (res []int) {
	next := solver(high)
	for num := next(low); num != 0; num = next(num + 1) {
		res = append(res, num)
	}
	return res
}

func main() {}
