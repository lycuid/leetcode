// https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/
package main

import "math"

type Graph struct{ parent []int }

func makeGraph(n int) Graph {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return Graph{parent}
}

func (g *Graph) union(x, y int) {
	if px, py := g.find(x), g.find(y); px != py {
		if px > py {
			g.parent[py] = px
		} else {
			g.parent[px] = py
		}
	}
}

func (g *Graph) find(x int) int {
	if x != g.parent[x] {
		g.parent[x] = g.find(g.parent[x])
	}
	return g.parent[x]
}

func minScore(n int, roads [][]int) int {
	graph := makeGraph(n + 1)
	cost := make([]int, n+1)
	for i := range cost {
		cost[i] = math.MaxInt
	}
	for _, edge := range roads {
		graph.union(edge[0], edge[1])
		cost[edge[0]] = min(cost[edge[0]], edge[2])
		cost[edge[1]] = min(cost[edge[1]], edge[2])
	}
	for i := 0; i < n; i++ {
		if graph.find(i) == n {
			cost[n] = min(cost[n], cost[i])
		}
	}
	return cost[n]
}

func main() {}
