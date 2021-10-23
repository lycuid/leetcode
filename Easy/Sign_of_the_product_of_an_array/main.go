// https://leetcode.com/problems/sign-of-the-product-of-an-array/
package main

func signum(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

func product(xs []int) int {
	prod := 1
	for _, x := range xs {
		prod *= x
	}
	return prod
}

func arraySign(xs []int) int {
	for i, x := range xs {
		xs[i] = signum(x)
	}
	return product(xs)
}

func main() {}
