// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/
package main

func kLengthApart(nums []int, k int) bool {
	count := k
	for _, num := range nums {
		if num == 1 {
			if count < k {
				return false
			}
			count = 0
		} else {
			count++
		}
	}
	return true
}

func main() {}
