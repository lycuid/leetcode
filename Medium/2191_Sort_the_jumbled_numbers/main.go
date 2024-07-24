// https://leetcode.com/problems/sort-the-jumbled-numbers/
package main

import "sort"

func sortJumbled(mapping []int, nums []int) []int {
	Convert := func(input int) (num int) {
		if input == 0 {
			return mapping[0]
		}
		for i := 1; input > 0; i *= 10 {
			num += mapping[input%10] * i
			input /= 10
		}
		return num
	}
	cache, ret := make([][2]int, len(nums)), make([]int, len(nums))
	for i, num := range nums {
		cache[i] = [2]int{Convert(num), i}
	}
	sort.Slice(cache, func(i, j int) bool {
		if cache[i][0] == cache[j][0] {
			return cache[i][1] < cache[j][1]
		}
		return cache[i][0] < cache[j][0]
	})
	for i, tup := range cache {
		ret[i] = nums[tup[1]]
	}
	return ret
}

func main() {}
