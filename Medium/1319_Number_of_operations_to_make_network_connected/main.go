// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
package main

type Graph struct{ inner []int }

func MakeGraph(n int) Graph {
	inner := make([]int, n)
	for i := range inner {
		inner[i] = i
	}
	return Graph{inner}
}

func (this *Graph) Union(x, y int) {
	if px, py := this.Find(x), this.Find(y); px != py {
		this.inner[py] = px
	}
}

func (this *Graph) Find(x int) int {
	if x != this.inner[x] {
		this.inner[x] = this.Find(this.inner[x])
	}
	return this.inner[x]
}

func (this *Graph) DisjointSets() (count int) {
	seen := make([]bool, len(this.inner))
	for i := 0; i < len(this.inner); i++ {
		if p := this.Find(i); !seen[p] {
			count, seen[p] = count+1, true
		}
	}
	return count
}

func makeConnected(n int, connections [][]int) int {
	graph, extra_edges := MakeGraph(n), 0
	for _, conn := range connections {
		if graph.Find(conn[0]) == graph.Find(conn[1]) {
			extra_edges++
		}
		graph.Union(conn[0], conn[1])
	}
	if count := graph.DisjointSets() - 1; count <= extra_edges {
		return count
	}
	return -1
}

func main() {}
