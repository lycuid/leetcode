// https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/
package main

const INF = 1e9 + 1

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Aux(ns, ms []int, left, i int, cache [][]int) int {
	if i == len(ms) {
		return 0
	}
	if cache[left][i] == INF {
		cache[left][i] = Max(
			(ms[i]*ns[left])+Aux(ns, ms, left+1, i+1, cache),
			(ms[i]*ns[len(ns)-1-i+left])+Aux(ns, ms, left, i+1, cache))
	}
	return cache[left][i]
}

func maximumScore(nums []int, multipliers []int) int {
	cache := make([][]int, len(multipliers))
	for i := range cache {
		cache[i] = make([]int, len(multipliers))
		for j := range cache[i] {
			cache[i][j] = INF
		}
	}
	return Aux(nums, multipliers, 0, 0, cache)
}

func main() {}
