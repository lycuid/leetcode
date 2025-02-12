// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
package main

func maximumSum(nums []int) int {
	digits, sum := make(map[int]int), -1
	dsum := func(n int) (sum int) {
		for ; n > 0; n /= 10 {
			sum += n % 10
		}
		return sum
	}
	for _, num := range nums {
		key := dsum(num)
		if m, found := digits[key]; found {
			sum = max(sum, m+num)
		}
		digits[key] = max(digits[key], num)
	}
	return sum
}

func main() {}
