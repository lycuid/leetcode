// https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/
package main

import "sort"

type Graph struct{ parent []int }

func MakeGraph(n int) Graph {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return Graph{parent}
}

func (this *Graph) Find(x int) int {
	if x != this.parent[x] {
		this.parent[x] = this.Find(this.parent[x])
	}
	return this.parent[x]
}

func (this *Graph) Union(x, y int) {
	if px, py := this.Find(x), this.Find(y); px != py {
		this.parent[py] = px
	}
}

func distanceLimitedPathsExist(n int, edges [][]int, queries [][]int) []bool {
	graph, ret := MakeGraph(n), make([]bool, len(queries))
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sorter := func(arr [][]int) func(int, int) bool {
		return func(i, j int) bool { return arr[i][2] < arr[j][2] }
	}
	sort.Slice(edges, sorter(edges))
	sort.Slice(queries, sorter(queries))
	for i, j := 0, 0; i < len(queries); i++ {
		for ; j < len(edges) && edges[j][2] < queries[i][2]; j++ {
			graph.Union(edges[j][0], edges[j][1])
		}
		ret[queries[i][3]] = graph.Find(queries[i][0]) == graph.Find(queries[i][1])
	}
	return ret
}

func main() {}
