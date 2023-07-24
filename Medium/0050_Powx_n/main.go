// https://leetcode.com/problems/powx-n/
package main

func Aux(x float64, n int) (pow float64) {
	if n == 0 {
		return 1
	}
	if pow = Aux(x, n/2); n%2 == 1 {
		return pow * pow * x
	}
	return pow * pow
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / Aux(x, -n)
	}
	return Aux(x, n)
}

func main() {}
