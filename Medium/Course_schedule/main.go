// https://leetcode.com/problems/course-schedule/
package main

func IsCyclic(root int, edges [][]int, callStack, visited *[2001]bool) bool {
	callStack[root], visited[root] = true, true
	for _, child := range edges[root] {
		if callStack[child] || (!visited[child] && IsCyclic(child, edges, callStack, visited)) {
			return true
		}
	}
	callStack[root] = false
	return false
}

func canFinish(n int, adj [][]int) bool {
	edges := make([][]int, n)
	for _, e := range adj {
		edges[e[0]] = append(edges[e[0]], e[1])
	}
	var callStack, visited [2001]bool
	for root := range edges {
		if !visited[root] && IsCyclic(root, edges, &callStack, &visited) {
			return false
		}
	}
	return true
}

func main() {}
