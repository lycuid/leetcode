// https://leetcode.com/problems/check-divisibility-by-digit-sum-and-product/
package main

func checkDivisibility(n int) bool {
	sum, prod := 0, 1
	for m := n; m > 0; m /= 10 {
		digit := m % 10
		sum, prod = sum+digit, prod*digit
	}
	return n%(sum+prod) == 0
}

func main() {}
