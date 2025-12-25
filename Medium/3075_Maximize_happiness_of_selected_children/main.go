// https://leetcode.com/problems/maximize-happiness-of-selected-children/
package main

import "sort"

func maximumHappinessSum(happiness []int, k int) (count int64) {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	for i, num := range happiness[:k] {
		count += int64(max(0, num-i))
	}
	return count
}

func main() {}
