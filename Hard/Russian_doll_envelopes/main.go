// https://leetcode.com/problems/russian-doll-envelopes/
package main

import "sort"

func maxEnvelopes(es [][]int) int {
	sort.Slice(es, func(i, j int) bool {
		if es[i][0] == es[j][0] {
			return es[i][1] > es[j][1]
		}
		return es[i][0] < es[j][0]
	})
	var lis []int
	lis = append(lis, es[0][1])
	for i := 1; i < len(es); i++ {
		if es[i][1] > lis[len(lis)-1] {
			lis = append(lis, es[i][1])
		} else {
			l, r := 0, len(lis)-1
			for l < r {
				mid := (l + r) / 2
				if lis[mid] < es[i][1] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			lis[l] = es[i][1]
		}
	}
	return len(lis)
}

func main() {}
