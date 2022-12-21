// https://leetcode.com/problems/possible-bipartition/
package main

func IsCyclic(root int, adj [][]int, id []int, vis []bool) bool {
	if !vis[root] {
		vis[root] = true
		for _, child := range adj[root] {
			if id[child] != 0 {
				if (id[root]-id[child]+1)%2 == 0 {
					continue
				}
				return true
			}
			if id[child] = id[root] + 1; IsCyclic(child, adj, id, vis) {
				return true
			}
			id[child] = 0
		}
	}
	return false
}

func possibleBipartition(n int, dislikes [][]int) bool {
	adj, vis := make([][]int, n+1), make([]bool, n+1)
	for _, edge := range dislikes {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	for i, id := 1, make([]int, n+1); i <= n; i++ {
		if id[i] = 1; IsCyclic(i, adj, id, vis) {
			return false
		}
		id[i] = 0
	}
	return true
}

func main() {}
