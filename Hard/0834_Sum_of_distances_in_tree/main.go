// https://leetcode.com/problems/sum-of-distances-in-tree/
package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	adj := make(map[int][]int)
	count, ret := make([]int, n), make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	var dfs, dfs1 func(int, int)
	dfs = func(node, parent int) {
		count[node] = 1
		for _, child := range adj[node] {
			if child != parent {
				dfs(child, node)
				count[node] += count[child]
				ret[node] += ret[child] + count[child]
			}
		}
	}

	dfs1 = func(node, parent int) {
		for _, child := range adj[node] {
			if child != parent {
				ret[child] = ret[node] + n - 2*count[child]
				dfs1(child, node)
			}
		}
	}
	dfs(0, -1)
	dfs1(0, -1)

	return ret
}

func main() {}
