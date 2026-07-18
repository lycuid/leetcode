// https://leetcode.com/problems/find-greatest-common-divisor-of-array/
package main

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findGCD(nums []int) int {
	low, high := nums[0], nums[0]
	for _, num := range nums {
		low, high = min(low, num), max(high, num)
	}
	return gcd(high, low)
}

func main() {}
