// https://leetcode.com/problems/unique-paths-ii/
package main

func dfs(x, y int, unique_path [][]int) {
	if unique_path[x][y] < 0 {
		unique_path[x][y] = 0
		if x < len(unique_path)-1 {
			dfs(x+1, y, unique_path)
			unique_path[x][y] += unique_path[x+1][y]
		}
		if y < len(unique_path[0])-1 {
			dfs(x, y+1, unique_path)
			unique_path[x][y] += unique_path[x][y+1]
		}
	}
}

func uniquePathsWithObstacles(graph [][]int) int {
	n, m := len(graph), len(graph[0])
	if graph[0][0]+graph[n-1][m-1] >= 1 {
		return 0
	}
	unique_path := make([][]int, n)
	for i := range unique_path {
		unique_path[i] = make([]int, m)
		for j := range unique_path[i] {
			unique_path[i][j] = graph[i][j] - 1
		}
	}
	unique_path[n-1][m-1] = 1
	dfs(0, 0, unique_path)
	return unique_path[0][0]
}

func main() {}
