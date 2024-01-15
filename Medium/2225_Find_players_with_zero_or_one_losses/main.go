// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
package main

import "sort"

type Tuple struct{ w, l int }

func findWinners(matches [][]int) [][]int {
	cache, ret := make(map[int]*Tuple), make([][]int, 2)
	for _, m := range matches {
		if player, found := cache[m[0]]; found {
			player.w++
		} else {
			cache[m[0]] = &Tuple{w: 1}
		}
		if player, found := cache[m[1]]; found {
			player.l++
		} else {
			cache[m[1]] = &Tuple{l: 1}
		}
	}
	for i, player := range cache {
		if loss := player.l; loss < 2 {
			ret[loss] = append(ret[loss], i)
		}
	}
	sort.Ints(ret[0])
	sort.Ints(ret[1])
	return ret
}

func main() {}
