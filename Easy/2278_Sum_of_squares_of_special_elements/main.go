// https://leetcode.com/problems/sum-of-squares-of-special-elements/
package main

func sumOfSquares(nums []int) (res int) {
	for i := range nums {
		if len(nums)%(i+1) == 0 {
			res += nums[i] * nums[i]
		}
	}
	return res
}

func main() {}
