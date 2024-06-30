// https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
package main

type Graph struct{ group []int }

func MakeGraph(n int) *Graph {
	group := make([]int, n)
	for i := range group {
		group[i] = i
	}
	return &Graph{group}
}

func (graph *Graph) Union(x, y int) {
	if gx, gy := graph.Find(x), graph.Find(y); gx != gy {
		graph.group[gy] = gx
	}
}

func (graph *Graph) Find(x int) int {
	if x != graph.group[x] {
		graph.group[x] = graph.Find(graph.group[x])
	}
	return graph.group[x]
}

func maxNumEdgesToRemove(n int, edges [][]int) (count int) {
	graphs := []*Graph{MakeGraph(n + 1), MakeGraph(n + 1)}
	for _, e := range edges {
		if e[0] == 3 {
			if graphs[0].Find(e[1]) == graphs[0].Find(e[2]) {
				count++
			} else {
				graphs[0].Union(e[1], e[2])
			}
		}
	}
	copy(graphs[1].group, graphs[0].group)
	for _, e := range edges {
		if e[0] != 3 {
			if graph := graphs[e[0]-1]; graph.Find(e[1]) == graph.Find(e[2]) {
				count++
			} else {
				graph.Union(e[1], e[2])
			}
		}
	}
	for i := 2; i <= n; i++ {
		for j := 0; j < len(graphs); j++ {
			if graphs[j].Find(i) != graphs[j].Find(i-1) {
				return -1
			}
		}
	}
	return count
}

func main() {}
