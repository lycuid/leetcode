// https://leetcode.com/problems/smallest-index-with-digit-sum-equal-to-index/
package main

func smallestIndex(nums []int) int {
	for i := range nums {
		var sum int
		for num := nums[i]; num > 0; num /= 10 {
			sum += num % 10
		}
		if sum == i {
			return i
		}
	}
	return -1
}

func main() {}
