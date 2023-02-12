// https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/
package main

func Dfs(root int, n int64, path []int64, adj map[int][]int, ret *int64) int64 {
	var rpath int64
	for _, node := range adj[root] {
		if path[node] == -1 {
			path[node] = path[root] + 1
			rpath += Dfs(node, n, path, adj, ret) - path[root]
		}
	}
	*ret += (path[root] + n) * (rpath / n)
	return path[root] + (rpath % n)
}

func minimumFuelCost(roads [][]int, seats int) (ret int64) {
	adj := make(map[int][]int)
	for _, road := range roads {
		adj[road[0]] = append(adj[road[0]], road[1])
		adj[road[1]] = append(adj[road[1]], road[0])
	}
	path := make([]int64, len(roads)+1)
	for i := 1; i < len(path); i++ {
		path[i] = -1
	}
	rpath := Dfs(0, int64(seats), path, adj, &ret)
	return ret + rpath
}

func main() {}
