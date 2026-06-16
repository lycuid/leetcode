// https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/
package main

import "sort"

func minimumCost(cost []int) (total int) {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})
	for i := range cost {
		if (i+1)%3 != 0 {
			total += cost[i]
		}
	}
	return total
}

func main() {}
