// https://leetcode.com/problems/count-the-number-of-complete-components/
package main

type Graph struct{ parent []int }

func makeGraph(n int) Graph {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return Graph{parent}
}

func (g *Graph) union(u, v int) {
	if pu, pv := g.find(u), g.find(v); pu != pv {
		g.parent[pv] = pu
	}
}

func (g *Graph) find(u int) int {
	if u != g.parent[u] {
		g.parent[u] = g.find(g.parent[u])
	}
	return g.parent[u]
}

func countCompleteComponents(n int, edges [][]int) int {
	var (
		graph   = makeGraph(n)
		edgecnt = make([]int, n)
		nodecnt = make(map[int]int)
	)
	for _, e := range edges {
		graph.union(e[0], e[1])
	}
	for _, e := range edges {
		edgecnt[graph.find(e[0])]++
	}
	for i := 0; i < n; i++ {
		nodecnt[graph.find(i)]++
	}
	for leader, count := range nodecnt {
		if edgecnt[leader] != (count*(count-1))/2 {
			delete(nodecnt, leader)
		}
	}
	return len(nodecnt)
}

func main() {}
