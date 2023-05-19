// https://leetcode.com/problems/is-graph-bipartite/
package main

func isBipartite(graph [][]int) bool {
	queue, set := []int{}, make([]int, len(graph))
	for parent := range graph {
		if set[parent] == 0 {
			set[parent], queue = 1, append(queue, parent)
			for root := 0; len(queue) > 0; {
				root, queue = queue[0], queue[1:]
				for _, child := range graph[root] {
					if set[child] == set[root] {
						return false
					}
					if set[child] == 0 {
						queue, set[child] = append(queue, child), 3-set[root]
					}
				}
			}
		}
	}
	return true
}

func main() {}
