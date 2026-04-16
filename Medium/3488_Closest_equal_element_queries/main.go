// https://leetcode.com/problems/closest-equal-element-queries/
package main

import "sort"

func solveQueries(nums []int, queries []int) []int {
	cache := make(map[int][]int)
	for i, num := range nums {
		cache[num] = append(cache[num], i)
	}
	dx := func(sub []int, x, target int) int {
		index := sub[(len(sub)+x)%len(sub)]
		if index > target {
			return min(index-target, len(nums)+target-index)
		}
		return min(target-index, len(nums)+index-target)
	}
	for i, q := range queries {
		queries[i] = -1
		sub, found := cache[nums[q]]
		if !found || len(sub) <= 1 {
			continue
		}
		index := sort.Search(len(sub), func(x int) bool {
			return sub[x] >= q
		})
		if index == len(sub) {
			continue
		}
		queries[i] = min(dx(sub, index+1, sub[index]), dx(sub, index-1, sub[index]))
	}
	return queries
}

func main() {}
