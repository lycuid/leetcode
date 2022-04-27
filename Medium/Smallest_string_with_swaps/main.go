// https://leetcode.com/problems/smallest-string-with-swaps/solution/
package main

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	p := make([]int, len(s))
	for i := range p {
		p[i] = i
	}
	find := func(x int) int {
		for x != p[x] {
			x, p[x] = p[x], p[p[x]]
		}
		return x
	}
	for _, pair := range pairs {
		p[find(pair[1])] = find(pair[0])
	}

	conns := make(map[int][]int)
	for i, index := range p {
		parent := find(index)
		conns[parent] = append(conns[parent], i)
	}

	ret := make([]byte, len(s))
	for _, conn := range conns {
		sub := make([]byte, len(conn))
		for parent, index := range conn {
			sub[parent] = s[index]
		}
		sort.Slice(sub, func(i, j int) bool { return sub[i] < sub[j] })
		for i, index := range conn {
			ret[index] = sub[i]
		}
	}
	return string(ret)
}

func main() {}
