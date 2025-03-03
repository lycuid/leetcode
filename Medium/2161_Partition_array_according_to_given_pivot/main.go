// https://leetcode.com/problems/partition-array-according-to-given-pivot/
package main

func pivotArray(nums []int, pivot int) []int {
	cache := make([]int, 0, len(nums))
	for _, num := range nums {
		if num < pivot {
			cache = append(cache, num)
		}
	}
	for _, num := range nums {
		if num == pivot {
			cache = append(cache, num)
		}
	}
	for _, num := range nums {
		if num > pivot {
			cache = append(cache, num)
		}
	}
	return cache
}

func main() {}
