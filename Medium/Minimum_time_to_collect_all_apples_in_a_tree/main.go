// https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
package main

func Aux(root int, adj [][]int, apples, visited []bool) (ret int) {
	visited[root] = true
	for _, node := range adj[root] {
		if !visited[node] {
			if n := Aux(node, adj, apples, visited); n > 0 || apples[node] {
				ret += (n + 2)
			}
		}
	}
	return ret
}

func minTime(n int, edges [][]int, hasApple []bool) (ret int) {
	adj := make([][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	return Aux(0, adj, hasApple, make([]bool, n))
}

func main() {}
