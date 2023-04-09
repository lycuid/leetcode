// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/
package main

func Max(n int, nums ...int) int {
	for i := range nums {
		if nums[i] > n {
			n = nums[i]
		}
	}
	return n
}

func Merge(a, b *[26]int) {
	for i := 0; i < 26; i++ {
		b[i] = Max(a[i], b[i])
	}
}

type Graph struct {
	visited, on_stack []bool
	cache             [][26]int
}

func MakeGraph(nodes int) Graph {
	return Graph{make([]bool, nodes), make([]bool, nodes), make([][26]int, nodes)}
}

func (graph *Graph) Dfs(root int, colors string, adj map[int][]int) bool {
	if graph.on_stack[root] {
		return false
	}
	if graph.on_stack[root] = true; !graph.visited[root] {
		graph.visited[root] = true
		for _, child := range adj[root] {
			if !graph.Dfs(child, colors, adj) {
				return false
			}
			Merge(&graph.cache[child], &graph.cache[root])
		}
		graph.cache[root][colors[root]-'a']++
	}
	graph.on_stack[root] = false
	return true
}

func largestPathValue(colors string, edges [][]int) int {
	adj := make(map[int][]int)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	var count [26]int
	graph := MakeGraph(len(colors))
	for node := range colors {
		if !graph.visited[node] {
			if !graph.Dfs(node, colors, adj) {
				return -1
			}
			Merge(&graph.cache[node], &count)
		}
	}
	return Max(count[0], count[1:]...)
}

func main() {}
