// https://leetcode.com/problems/number-of-ways-to-assign-edge-weights-i/
package main

func assignEdgeWeights(edges [][]int) (res int) {
	adj := make(map[int][]int)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	var (
		depth   int
		cache   = make([]int, len(edges)+2)
		visited = make([]bool, len(edges)+2)
		queue   = make([]int, 0, len(edges)+2)
	)

	for queue = append(queue, 1); len(queue) > 0; queue = queue[1:] {
		node := queue[0]
		visited[node] = true
		for _, child := range adj[node] {
			if !visited[child] {
				queue = append(queue, child)
				cache[child] = 1 + cache[node]
				depth = max(depth, cache[child])
			}
		}
	}
	for res = 1; depth > 1; depth -= 31 {
		res = (res << min(depth-1, 31)) % int(1e9+7)
	}
	return res
}

func main() {}
