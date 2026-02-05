// https://leetcode.com/problems/transformed-array/
package main

func constructTransformedArray(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		for num = i + num; num < 0; {
			num += len(nums)
		}
		result[i] = nums[num%len(nums)]
	}
	return result
}

func main() {}
