// https://leetcode.com/problems/maximum-walls-destroyed-by-robots/
package main

import "sort"

func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	rs := func(arr []int, left, right int) int {
		if left > right {
			return 0
		}
		l := sort.Search(len(arr), func(i int) bool {
			return arr[i] >= left
		})
		r := sort.Search(len(arr), func(i int) bool {
			return arr[i] > right
		})
		return r - l
	}
	sort.Ints(walls)

	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return robots[idx[i]] < robots[idx[j]]
	})
	cache := [2][2]int{{0, 0}, {rs(walls, robots[idx[0]]-distance[idx[0]], robots[idx[0]]), 0}}
	for i := 1; i < n; i++ {
		cr, cd := robots[idx[i]], distance[idx[i]]
		pr, pd := robots[idx[i-1]], distance[idx[i-1]]
		cleft := rs(walls, max(pr+1, cr-cd), cr)
		pright := rs(walls, pr, min(pr+pd, cr-1))
		cache[1][1] += pright
		cache[0], cache[1] = cache[1], cache[0]

		cache[1][0] = max(cache[0][0]+cleft, cache[0][1]+cleft-rs(walls, max(pr+1, cr-cd), min(pr+pd, cr-1)))
		cache[1][1] = max(cache[0][0], cache[0][1])
	}
	cache[1][1] += rs(walls, robots[idx[n-1]], robots[idx[n-1]]+distance[idx[n-1]])
	return max(cache[1][0], cache[1][1])
}

func main() {}
