// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(i, j int) bool { return potions[i] > potions[j] })
	for i, n := 0, len(potions); i < len(spells); i++ {
		value, r := spells[i], n-1
		for spells[i] = 0; spells[i] <= r; {
			if mid := (spells[i] + r) / 2; int64(potions[mid]*value) < success {
				r = mid - 1
			} else {
				spells[i] = mid + 1
			}
		}
	}
	return spells
}

func main() {}
