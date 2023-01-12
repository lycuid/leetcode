// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/
package main

func Aux(root int, adj [][]int, labels string, visited []bool, chars [][26]int) {
	visited[root], chars[root][labels[root]-'a'] = true, chars[root][labels[root]-'a']+1
	for _, node := range adj[root] {
		if !visited[node] {
			Aux(node, adj, labels, visited, chars)
			for ch, count := range chars[node] {
				chars[root][ch] += count
			}
		}
	}
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	adj, ret, chars := make([][]int, n), make([]int, n), make([][26]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	Aux(0, adj, labels, make([]bool, n), chars)
	for i := 0; i < n; i++ {
		ret[i] = chars[i][labels[i]-'a']
	}
	return ret
}

func main() {}
