// https://leetcode.com/problems/relative-ranks/
package main

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	cache, rank := make([]int, len(score)), make([]string, len(score))
	for i := range cache {
		cache[i] = i
	}
	sort.Slice(cache, func(i, j int) bool {
		return score[cache[i]] > score[cache[j]]
	})
	for i := range cache {
		switch i {
		case 0:
			rank[cache[i]] = "Gold Medal"
		case 1:
			rank[cache[i]] = "Silver Medal"
		case 2:
			rank[cache[i]] = "Bronze Medal"
		default:
			rank[cache[i]] = fmt.Sprintf("%d", i+1)
		}
	}
	return rank
}

func main() {}
