// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
package main

func minReorder(n int, connections [][]int) (count int) {
	adj, rev := make([][]int, n), make([][]int, n)
	for _, conn := range connections {
		adj[conn[0]] = append(adj[conn[0]], conn[1])
		rev[conn[1]] = append(rev[conn[1]], conn[0])
	}
	on_queue := make([]bool, n)
	on_queue[0] = true
	for queue := []int{0}; len(queue) > 0; queue = queue[1:] {
		root := queue[0]
		for _, child := range adj[root] {
			if !on_queue[child] {
				queue, on_queue[child], count = append(queue, child), true, count+1
			}
		}
		for _, child := range rev[root] {
			if !on_queue[child] {
				queue, on_queue[child] = append(queue, child), true
			}
		}
	}
	return count
}

func main() {}
