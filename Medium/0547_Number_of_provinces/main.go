// https://leetcode.com/problems/number-of-provinces/
package main

type Graph struct{ parent []int }

func MakeGraph(n int) *Graph {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &Graph{parent}
}

func (this *Graph) Union(x, y int) {
	if px, py := this.Find(x), this.Find(y); px != py {
		this.parent[py] = px
	}
}

func (this *Graph) Find(x int) int {
	if this.parent[x] != x {
		this.parent[x] = this.Find(this.parent[x])
	}
	return this.parent[x]
}

func findCircleNum(isConnected [][]int) int {
	graph, groups := MakeGraph(len(isConnected)), make(map[int]int)
	for i := range isConnected {
		for j := range isConnected[i] {
			if isConnected[i][j] == 1 {
				graph.Union(i, j)
			}
		}
	}
	for i := range isConnected {
		groups[graph.Find(i)]++
	}
	return len(groups)
}

func main() {}
