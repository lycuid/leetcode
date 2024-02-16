// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/
package main

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	var cache []int
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}
	for _, value := range freq {
		cache = append(cache, value)
	}
	sort.Ints(cache)
	for k > 0 {
		if k -= cache[0]; k >= 0 {
			cache = cache[1:]
		}
	}
	return len(cache)
}

func main() {}
