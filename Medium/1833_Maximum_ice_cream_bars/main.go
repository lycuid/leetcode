// https://leetcode.com/problems/maximum-ice-cream-bars/
package main

import "sort"

func maxIceCream(costs []int, coins int) (i int) {
	sort.Ints(costs)
	for i, costs = 1, append([]int{0}, costs...); i < len(costs); i++ {
		if costs[i] += costs[i-1]; costs[i] > coins {
			break
		}
	}
	return i - 1
}

func main() {}
