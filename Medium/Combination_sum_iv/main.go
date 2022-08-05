// https://leetcode.com/problems/combination-sum-iv/
package main

func combinationSum4(nums []int, target int) int {
	cache := make([]int, target+1)
	cache[0] = 1
	for t := 1; t <= target; t++ {
		for _, num := range nums {
			if num <= t {
				cache[t] += cache[t-num]
			}
		}
	}
	return cache[target]
}

func main() {}
