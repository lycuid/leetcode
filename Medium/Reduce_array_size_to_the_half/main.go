// https://leetcode.com/problems/reduce-array-size-to-the-half/
package main

import "sort"

func minSetSize(arr []int) int {
	var freq [1e5 + 1]int
	var nums []int
	for _, num := range arr {
		if freq[num]++; freq[num] == 1 {
			nums = append(nums, num)
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return freq[nums[i]] > freq[nums[j]]
	})
	for i, sum := 0, 0; i < len(nums); i++ {
		if sum += freq[nums[i]]; sum >= len(arr)/2 {
			return i + 1
		}
	}
	return 0
}

func main() {}
