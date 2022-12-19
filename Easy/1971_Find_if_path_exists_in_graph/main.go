// https://leetcode.com/problems/find-if-path-exists-in-graph/
package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	adj := make(map[int][]int)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	for bfs, done, root := []int{source}, map[int]bool{source: true}, 0; len(bfs) > 0; {
		if root, bfs = bfs[0], bfs[1:]; root == destination {
			return true
		}
		for _, child := range adj[root] {
			if !done[child] {
				done[child], bfs = true, append(bfs, child)
			}
		}
	}
	return false
}

func main() {}
