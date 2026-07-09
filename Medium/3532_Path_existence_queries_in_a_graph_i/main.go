// https://leetcode.com/problems/path-existence-queries-in-a-graph-i/
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
		g.parent[pu] = pv
	}
}

func (g *Graph) find(u int) int {
	if u != g.parent[u] {
		g.parent[u] = g.find(g.parent[u])
	}
	return g.parent[u]
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	var (
		res   = make([]bool, len(queries))
		graph = makeGraph(n)
	)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] <= maxDiff {
			graph.union(i, i-1)
		}
	}
	for i, q := range queries {
		res[i] = graph.find(q[0]) == graph.find(q[1])
	}
	return res
}

func main() {}
