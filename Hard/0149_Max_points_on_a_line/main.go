// https://leetcode.com/problems/max-points-on-a-line/
package main

import "sort"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxPoints(ps [][]int) (ret int) {
	sort.Slice(ps, func(i, j int) bool { return ps[i][0] < ps[j][0] })
	cache := make([][]int, len(ps))
	for i := range cache {
		cache[i] = make([]int, len(ps))
	}
	for i := 0; i < len(ps); i++ {
		for j := i; j < len(ps); j++ {
			for k := j; k >= i; k-- {
				if (ps[k][1]-ps[i][1])*(ps[j][0]-ps[k][0]) == (ps[j][1]-ps[k][1])*(ps[k][0]-ps[i][0]) {
					cache[i][j] = Max(cache[i][j], 1+cache[i][k])
				}
			}
			ret = Max(ret, cache[i][j])
		}
	}
	return ret
}

func main() {}
