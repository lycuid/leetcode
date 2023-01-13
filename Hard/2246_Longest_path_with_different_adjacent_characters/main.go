// https://leetcode.com/problems/longest-path-with-different-adjacent-characters/
package main

func Dfs(root int, s string, edge []map[int]int) int {
	ret := 1
	for node := range edge[root] {
		n := Dfs(node, s, edge)
		if edge[root][node] = 1; s[root] != s[node] {
			edge[root][node] += n
		}
		if edge[root][node] > ret {
			ret = edge[root][node]
		}
	}
	return ret
}

func longestPath(parent []int, s string) (ret int) {
	edge := make([]map[int]int, len(parent))
	for i := 1; i < len(parent); i++ {
		if edge[parent[i]] == nil {
			edge[parent[i]] = make(map[int]int)
		}
		edge[parent[i]][i] = 0
	}
	Dfs(0, s, edge)
	for root := range edge {
		path1, path2 := 1, 1
		for node := range edge[root] {
			if edge[root][node] > path1 {
				path1, path2 = edge[root][node], path1
			} else if edge[root][node] > path2 {
				path2 = edge[root][node]
			}
		}
		if ret < path1+path2-1 {
			ret = path1 + path2 - 1
		}
	}
	return ret
}

func main() {}
