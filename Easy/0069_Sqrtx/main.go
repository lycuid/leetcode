// https://leetcode.com/problems/sqrtx/
package main

func mySqrt(x int) (l int) {
	for r := x; l < r; {
		if mid := (l + r) / 2; mid*mid > x {
			r = mid
		} else if mid*mid < x {
			l = mid + 1
		} else {
			l, r = mid, mid
		}
	}
	for l*l > x {
		l--
	}
	return l
}

func main() {}
