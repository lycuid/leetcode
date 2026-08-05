// https://leetcode.com/problems/remove-methods-from-project/
package main

func remainingMethods(n int, k int, invocations [][]int) []int {
	var (
		res     = make([]int, 0, n)
		visited = make([]bool, n)
		sus     = make([]bool, n)
		adj     = make(map[int][]int)
	)
	for _, e := range invocations {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	stack := make([]int, 0, n)
	for stack = append(stack, k); len(stack) > 0; stack = stack[1:] {
		node := stack[0]
		if sus[node] {
			continue
		}
		sus[node] = true
		for _, child := range adj[node] {
			stack = append(stack, child)
		}
	}
	stack = stack[:0]
	for i := 0; i < n; i++ {
		if !sus[i] && !visited[i] {
			stack, visited[i] = append(stack, i), true
		}
		for ; len(stack) > 0; stack = stack[1:] {
			node := stack[0]
			if sus[node] {
				goto all
			}
			for _, child := range adj[node] {
				if !visited[child] {
					stack, visited[child] = append(stack, child), true
				}
			}
		}
	}
	for i := range sus {
		if !sus[i] {
			res = append(res, i)
		}
	}
	return res
all:
	for i := 0; i < n; i++ {
		res = append(res, i)
	}
	return res
}

func main() {}
