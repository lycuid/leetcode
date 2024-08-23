// https://leetcode.com/problems/fraction-addition-and-subtraction/
package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return abs(a) * abs(b) / gcd(a, b)
}

func ParseNumber(expr string) (int, string) {
	negative, num := false, 0
	if expr[0] == '-' || expr[0] == '+' {
		negative, expr = expr[0] == '-', expr[1:]
	}
	for ; len(expr) > 0 && expr[0] >= '0' && expr[0] <= '9'; expr = expr[1:] {
		num = num*10 + int(expr[0]-'0')
	}
	if negative {
		num = -num
	}
	return num, expr
}

type Fraction struct{ n, d int }

func MakeFraction(expr string) (frac Fraction, _ string) {
	frac.n, expr = ParseNumber(expr)
	expr = expr[1:]
	frac.d, expr = ParseNumber(expr)
	return frac, expr
}

func (frac *Fraction) ToReducedForm() {
	for div := gcd(abs(frac.n), abs(frac.d)); div > 1; div = gcd(abs(frac.n), abs(frac.d)) {
		frac.n, frac.d = frac.n/div, frac.d/div
	}
}

func (frac *Fraction) Mul(mul int) {
	frac.n, frac.d = frac.n*mul, frac.d*mul
}

func (this *Fraction) Add(that Fraction) {
	mul := lcm(this.d, that.d)
	this.Mul(mul / this.d)
	that.Mul(mul / that.d)
	this.n += that.n
	this.ToReducedForm()
}

func fractionAddition(expression string) string {
	result := Fraction{0, 1}
	for len(expression) > 0 {
		frac, expr := MakeFraction(expression)
		result.Add(frac)
		expression = expr
	}
	return fmt.Sprintf("%d/%d", result.n, result.d)
}

func main() {}
