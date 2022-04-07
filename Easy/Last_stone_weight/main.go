// https://leetcode.com/problems/last-stone-weight/
package main

import "sort"

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	i := len(stones) - 2
	for ; i >= 0; i-- {
		if stones[i] == stones[i+1] {
			i--
			continue
		}
		stones[i] = stones[i+1] - stones[i]
		for j, k := i-1, i; j >= 0; j-- {
			if stones[j] <= stones[k] {
				break
			}
			stones[j], stones[k] = stones[k], stones[j]
			k--
		}
	}
	if i < -1 {
		return 0
	}
	return stones[0]
}

func main() {}
