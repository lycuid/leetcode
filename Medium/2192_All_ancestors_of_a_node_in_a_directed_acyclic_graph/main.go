// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/
package main

func getAncestors(n int, edges [][]int) [][]int {
	var (
		adj       = make([][]int, n)
		ancestors = make([][]int, n)
		Dfs       func(int)
	)

	for _, e := range edges {
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	Dfs = func(node int) {
		if len(ancestors[node]) == 0 {
			cache := make([]bool, n)
			for _, child := range adj[node] {
				cache[child] = true
				Dfs(child)
				for _, anc := range ancestors[child] {
					cache[anc] = true
				}
			}
			for anc, found := range cache {
				if found {
					ancestors[node] = append(ancestors[node], anc)
				}
			}
		}
	}

	for node := 0; node < n; node++ {
		Dfs(node)
	}

	return ancestors
}

func main() {}
