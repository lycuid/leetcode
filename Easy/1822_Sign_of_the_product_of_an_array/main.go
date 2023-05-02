// https://leetcode.com/problems/sign-of-the-product-of-an-array/
package main

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func arraySign(nums []int) int {
	ret := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		ret *= (num / Abs(num))
	}
	return ret
}

func main() {}
