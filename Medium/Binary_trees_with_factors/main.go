// https://leetcode.com/problems/binary-trees-with-factors/
package main

import "sort"

const MOD = 1e9 + 7

func numFactoredBinaryTrees(nums []int) (ret int) {
	cache := make(map[int]int)
	sort.Ints(nums)
	for i := range nums {
		cache[nums[i]] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && cache[nums[i]/nums[j]] > 0 {
				cache[nums[i]] += (cache[nums[j]] * cache[nums[i]/nums[j]])
			}
		}
		ret = (ret + cache[nums[i]]) % MOD
	}
	return ret % MOD
}

func main() {}
