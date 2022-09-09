// https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
package main

import "sort"

func numberOfWeakCharacters(p [][]int) (ret int) {
	sort.Slice(p, func(i, j int) bool {
		if p[i][0] == p[j][0] {
			return p[i][1] < p[j][1]
		}
		return p[i][0] > p[j][0]
	})
	var min int
	for i := range p {
		if p[i][1] < min {
			ret++
		} else {
			min = p[i][1]
		}
	}
	return ret
}

func main() {}
