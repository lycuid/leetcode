// https://leetcode.com/problems/count-symmetric-integers/
package main

func countSymmetricIntegers(low int, high int) (count int) {
	var cache [10]int
	Digits := func(num int) []int {
		var i int
		for ; num > 0; num /= 10 {
			cache[i], i = num%10, i+1
		}
		return cache[:i]
	}
	Sum := func(nums []int) (num int) {
		for i := range nums {
			num += nums[i]
		}
		return num
	}
	for num := low; num <= high; num++ {
		digits := Digits(num)
		if n := len(digits); n%2 == 0 {
			if mid := n / 2; Sum(digits[:mid]) == Sum(digits[mid:]) {
				count++
			}
		}
	}
	return count
}

func main() {}
