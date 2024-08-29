// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
package main

type Graph struct{ parent []int }

func MakeGraph(n int) Graph {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return Graph{parent}
}

func (graph *Graph) Union(x, y int) {
	if px, py := graph.Find(x), graph.Find(y); px != py {
		graph.parent[py] = px
	}
}

func (graph *Graph) Find(x int) int {
	if graph.parent[x] != x {
		graph.parent[x] = graph.Find(graph.parent[x])
	}
	return graph.parent[x]
}

func removeStones(stones [][]int) (count int) {
	AreGrouped := func(i, j int) bool {
		return stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1]
	}
	graph := MakeGraph(len(stones))
	for i := range stones {
		for j := i + 1; j < len(stones); j++ {
			if AreGrouped(i, j) {
				graph.Union(i, j)
			}
		}
	}
	groups := make([]bool, len(stones))
	for i := range stones {
		if g := graph.Find(i); !groups[g] {
			groups[g], count = true, count+1
		}
	}
	return len(stones) - count
}

func main() {}
