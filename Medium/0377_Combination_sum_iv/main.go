// https://leetcode.com/problems/combination-sum-iv/
package main

func combinationSum4(nums []int, target int) int {
	cache := make([]int, target+1)
	cache[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				cache[i] += cache[i-num]
			}
		}
	}
	return cache[target]
}

func main() {}
