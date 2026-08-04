// https://leetcode.com/problems/find-missing-elements/
package main

func findMissingElements(nums []int) (res []int) {
	low, high := nums[0], nums[0]
	for _, num := range nums {
		low, high = min(low, num), max(high, num)
	}
	cache := make([]bool, high-low+1)
	for _, num := range nums {
		cache[num-low] = true
	}
	for i := range cache {
		if !cache[i] {
			res = append(res, low+i)
		}
	}
	return res
}

func main() {}
