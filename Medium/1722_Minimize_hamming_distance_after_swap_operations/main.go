// https://leetcode.com/problems/minimize-hamming-distance-after-swap-operations/
package main

type Graph struct{ parent []int }

func NewGraph(n int) *Graph {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &Graph{parent}
}

func (graph *Graph) union(x, y int) {
	if px, py := graph.find(x), graph.find(y); px != py {
		graph.parent[py] = px
	}
}

func (graph *Graph) find(node int) int {
	if graph.parent[node] != node {
		graph.parent[node] = graph.find(graph.parent[node])
	}
	return graph.parent[node]
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) (res int) {
	graph := NewGraph(len(source))
	for _, swap := range allowedSwaps {
		graph.union(swap[0], swap[1])
	}
	cache := make([]map[int]int, len(source))
	for i, val := range source {
		root := graph.find(i)
		if cache[root] == nil {
			cache[root] = make(map[int]int)
		}
		cache[root][val]++
	}
	for i, val := range target {
		root := graph.find(i)
		if cnt, found := cache[root][val]; found {
			if cnt == 1 {
				delete(cache[root], val)
			} else {
				cache[root][val] = cnt - 1
			}
		} else {
			res++
		}
	}
	return res
}

func main() {}
