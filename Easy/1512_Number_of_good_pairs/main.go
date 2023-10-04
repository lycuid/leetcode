// https://leetcode.com/problems/number-of-good-pairs/
package main

func numIdenticalPairs(nums []int) (ret int) {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}
	for _, count := range cache {
		ret += (count * (count - 1)) / 2
	}
	return ret
}

func main() {}
