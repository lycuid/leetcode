// https://leetcode.com/problems/is-graph-bipartite/
package main

const (
	Orphan = iota
	Group1 = iota
	Group2 = iota
)

func TryPartition(parent int, children [][]int, group_of []int) bool {
	new_group := Group1 + Group2 - group_of[parent]
	for _, child := range children[parent] {
		switch group_of[child] {
		case Orphan:
			group_of[child] = new_group
		case group_of[parent]:
			return false
		default:
			continue
		}
		if !TryPartition(child, children, group_of) {
			return false
		}
	}
	return true
}

func isBipartite(graph [][]int) bool {
	group_of := make([]int, len(graph)) // 1 <= len(graph) <= 100
	for node := range graph {
		if group_of[node] != Orphan {
			continue
		}
		group_of[node] = Group1
		if !TryPartition(node, graph, group_of) {
			return false
		}
	}
	return true
}

func main() {}
