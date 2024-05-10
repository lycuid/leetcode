// https://leetcode.com/problems/k-th-smallest-prime-fraction/
package main

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var cache [][]int
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			cache = append(cache, []int{arr[i], arr[j]})
		}
	}
	sort.Slice(cache, func(i, j int) bool {
		return float32(cache[i][0])/float32(cache[i][1]) < float32(cache[j][0])/float32(cache[j][1])
	})
	return cache[k-1]
}

func main() {}
