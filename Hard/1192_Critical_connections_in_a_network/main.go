// https://leetcode.com/problems/critical-connections-in-a-network/
package main

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type Network struct {
	graph   map[int][]int
	id, low []int
	newid   int
	visited []bool
	ret     [][]int
}

func Init(n int, edges [][]int) Network {
	graph := make(map[int][]int)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return Network{graph: graph, id: make([]int, n), low: make([]int, n), visited: make([]bool, n)}
}

func (this *Network) CriticalConnections() [][]int {
	this.dfs(0, -1)
	return this.ret
}

func (this *Network) dfs(node, parent int) {
	this.id[node], this.low[node] = this.newid, this.newid
	this.newid++
	this.visited[node] = true
	for _, child := range this.graph[node] {
		if child == parent {
			continue
		}
		if !this.visited[child] {
			this.dfs(child, node)
		}
		this.low[node] = Min(this.low[node], this.low[child])
		if this.id[node] < this.low[child] {
			this.ret = append(this.ret, []int{node, child})
		}
	}
}

func criticalConnections(n int, connections [][]int) [][]int {
	network := Init(n, connections)
	return network.CriticalConnections()
}

func main() {}
