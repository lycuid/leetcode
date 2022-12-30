// https://leetcode.com/problems/all-paths-from-source-to-target/
package main

func Dfs(i, target int, graph [][]int, cache *[][][]int) [][]int {
	if i == target {
		return [][]int{{target}}
	}
	if len((*cache)[i]) == 0 {
		for _, node := range graph[i] {
			for _, path := range Dfs(node, target, graph, cache) {
				(*cache)[i] = append((*cache)[i], append([]int{i}, path...))
			}
		}
	}
	return (*cache)[i]
}

func allPathsSourceTarget(graph [][]int) [][]int {
	cache := make([][][]int, len(graph))
	return Dfs(0, len(graph)-1, graph, &cache)
}

func main() {}
