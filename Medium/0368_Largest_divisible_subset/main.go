// https://leetcode.com/problems/largest-divisible-subset/
package main

import "sort"

func largestDivisibleSubset(nums []int) (ret []int) {
	sort.Ints(nums)

	cache, max := make([][2]int, len(nums)), 0
	for i := range cache {
		cache[i] = [2]int{i, 1}
	}
	for i := range cache {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && cache[j][1] >= cache[i][1] {
				if cache[i] = [2]int{j, cache[j][1] + 1}; cache[i][1] > cache[max][1] {
					max = i
				}
			}
		}
	}
	for ; max != cache[max][0]; max = cache[max][0] {
		ret = append(ret, nums[max])
	}
	ret = append(ret, nums[max])

	return ret
}

func main() {}
