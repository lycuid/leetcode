// https://leetcode.com/problems/smallest-divisible-digit-product-i/
package main

func smallestNumber(n int, t int) int {
	digit_prod := func(n int) (res int) {
		for res = 1; n > 0; n /= 10 {
			res *= n % 10
		}
		return res
	}
	for ; ; n++ {
		if prod := digit_prod(n); prod%t == 0 {
			break
		}
	}
	return n
}

func main() {}
