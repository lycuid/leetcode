// https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
package main

import (
	"fmt"
)

type Graph struct {
	groups int
	parent []int
}

func MakeGraph(n int) Graph {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return Graph{n, parent}
}

func (this *Graph) Union(x, y int) (unified int) {
	if px, py := this.Find(x), this.Find(y); px != py {
		this.parent[py], this.groups, unified = px, this.groups-1, 1
	}
	return unified
}

func (this *Graph) Find(x int) int {
	if x != this.parent[x] {
		this.parent[x] = this.Find(this.parent[x])
	}
	return this.parent[x]
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	graph, ret := [2]Graph{MakeGraph(n), MakeGraph(n)}, len(edges)
	for i, j := 0, len(edges)-1; i < j; i++ {
		if edges[i][0] != 3 {
			edges[i], edges[j], i, j = edges[j], edges[i], i-1, j-1
		}
	}
	for _, edge := range edges {
		if e, u, v := edge[0]-1, edge[1]-1, edge[2]-1; e < 2 {
			ret -= graph[e].Union(u, v)
		} else {
			ret -= graph[0].Union(u, v) | graph[1].Union(u, v)
		}
	}
	if graph[0].groups != 1 || graph[1].groups != 1 {
		return -1
	}
	return ret
}

func main() {
	fmt.Println(maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}))
	fmt.Println(maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}))
	fmt.Println(maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}}))
}
