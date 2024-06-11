// https://leetcode.com/problems/relative-sort-array/
package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	cache := make(map[int]int)
	for i, num := range arr2 {
		cache[num] = i + 1
	}
	sort.Slice(arr1, func(i, j int) bool {
		var num1, num2 int
		if num1 = cache[arr1[i]]; num1 == 0 {
			num1 = arr1[i] + len(arr2) + 1
		}
		if num2 = cache[arr1[j]]; num2 == 0 {
			num2 = arr1[j] + len(arr2) + 1
		}
		return num1 < num2
	})
	return arr1
}

func main() {}
