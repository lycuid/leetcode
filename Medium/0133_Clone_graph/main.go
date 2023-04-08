// https://leetcode.com/problems/clone-graph/
package main

type Node struct {
	Val       int
	Neighbors []*Node
}

type Graph struct{ cache map[*Node]*Node }

func (graph *Graph) Clone(node *Node) *Node {
	if node != nil && graph.cache[node] == nil {
		graph.cache[node] = &Node{Val: node.Val}
		for _, n := range node.Neighbors {
			graph.cache[node].Neighbors = append(
				graph.cache[node].Neighbors,
				graph.Clone(n),
			)
		}
	}
	return graph.cache[node]
}

func cloneGraph(node *Node) *Node {
	graph := Graph{make(map[*Node]*Node)}
	return graph.Clone(node)
}

func main() {}
