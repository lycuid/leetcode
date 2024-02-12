// https://leetcode.com/problems/majority-element/
package main

func majorityElement(nums []int) int {
	freq, max, num := make(map[int]int), 0, nums[0]
	for i := range nums {
		if freq[nums[i]]++; freq[nums[i]] > max {
			max, num = freq[nums[i]], nums[i]
		}
	}
	return num
}

func main() {}
