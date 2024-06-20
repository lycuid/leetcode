// https://leetcode.com/problems/magnetic-force-between-two-balls/
package main

import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	ret := -1
	for l, r := 1, position[len(position)-1]-position[0]; l <= r; {
		mid := l + (r-l)/2
		last_pos, balls := position[0], 1
		for _, pos := range position {
			if pos-last_pos >= mid {
				last_pos, balls = pos, balls+1
			}
		}
		if balls >= m {
			l, ret = mid+1, mid
		} else {
			r = mid - 1
		}
	}
	return ret
}

func main() {}
