// https://leetcode.com/problems/network-recovery-pathways/
package main

import (
	"math"
	"sort"
)

type Item struct {
	node   int
	weight int64
}

type PriorityQueue struct{ items []Item }

func makePQ() PriorityQueue {
	return PriorityQueue{make([]Item, 1)}
}

func (pq PriorityQueue) less(i, j int) bool {
	return pq.items[i].weight < pq.items[j].weight
}

func (pq PriorityQueue) empty() bool { return len(pq.items) == 1 }

func (pq *PriorityQueue) push(item Item) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.less(i, i/2); i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *PriorityQueue) pop() Item {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.less(j+1, j) {
			j++
		}
		if pq.less(i, j) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	var (
		n     = len(online)
		adj   = make([][][2]int, n)
		_cost = make([]int, 0, n)
	)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], [2]int{e[1], e[2]})
		_cost = append(_cost, e[2])
	}
	sort.Ints(_cost)
	cost := _cost[:0]
	for _, c := range _cost {
		if n := len(cost); n == 0 || cost[n-1] != c {
			cost = append(cost, c)
		}
	}

	check := func(limit int) bool {
		pq, edgemin := makePQ(), make([]int64, n)
		for i := range edgemin {
			edgemin[i] = math.MaxInt64
		}
		edgemin[0] = 0
		for pq.push(Item{0, 0}); !pq.empty(); {
			item := pq.pop()
			if item.weight > k || edgemin[item.node] != item.weight {
				continue
			}
			if item.node == len(online)-1 {
				return true
			}

			for _, child := range adj[item.node] {
				if child[1] < limit || (child[0] != n-1 && !online[child[0]]) {
					continue
				}
				if d := item.weight + int64(child[1]); d < edgemin[child[0]] {
					edgemin[child[0]] = d
					pq.push(Item{child[0], d})
				}
			}
		}
		return false
	}

	res := -1
	for l, r := 0, len(cost)-1; l <= r; {
		if mid := l + (r-l)/2; check(cost[mid]) {
			res = cost[mid]
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}

func main() {}
