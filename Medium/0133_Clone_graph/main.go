// https://leetcode.com/problems/clone-graph/
package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func Clone(node *Node, cache *[101]*Node) *Node {
	if node == nil {
		return nil
	}
	if cache[node.Val] != nil {
		return cache[node.Val]
	}
	cache[node.Val] = &Node{node.Val, make([]*Node, len(node.Neighbors))}
	for i, neighbour := range node.Neighbors {
		cache[node.Val].Neighbors[i] = Clone(neighbour, cache)
	}
	return cache[node.Val]
}

func cloneGraph(node *Node) *Node {
	var cache [101]*Node
	return Clone(node, &cache)
}

func main() {}
