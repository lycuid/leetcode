// https://leetcode.com/problems/minimum-height-trees/
package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj, deg := make(map[int][]int), make([]int, n)
	for _, e := range edges {
		adj[e[0]], adj[e[1]] = append(adj[e[0]], e[1]), append(adj[e[1]], e[0])
		deg[e[0]], deg[e[1]] = deg[e[0]]+1, deg[e[1]]+1
	}
	var queue []int
	for node, d := range deg {
		if d == 1 {
			queue = append(queue, node)
		}
	}
	for n > 2 {
		size := len(queue)
		n -= size
		for i := 0; i < size; i, queue = i+1, queue[1:] {
			leaf := queue[0]
			for _, child := range adj[leaf] {
				if deg[child]--; deg[child] == 1 {
					queue = append(queue, child)
				}
			}
		}
	}
	return queue
}

func main() {}
