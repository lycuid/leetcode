// https://leetcode.com/problems/sum-of-even-numbers-after-queries/
package main

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	ret := make([]int, len(nums)+1)
	for i := range nums {
		if nums[i]%2 == 0 {
			ret[0] += nums[i]
		}
	}
	for i, query := range queries {
		val, index := query[0], query[1]
		if ret[i+1] = ret[i]; index < len(nums) {
			if nums[index]%2 == 0 {
				ret[i+1] -= nums[index]
			}
			if nums[index] += val; nums[index]%2 == 0 {
				ret[i+1] += nums[index]
			}
		}
	}
	return ret[1:]
}

func main() {}
