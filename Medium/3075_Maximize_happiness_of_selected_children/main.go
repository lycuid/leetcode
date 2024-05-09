// https://leetcode.com/problems/maximize-happiness-of-selected-children/
package main

import "sort"

func maximumHappinessSum(happiness []int, k int) (count int64) {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	for i, num := range happiness[:k] {
		if n := int64(num - i); n > 0 {
			count += n
		}
	}
	return count
}

func main() {}
