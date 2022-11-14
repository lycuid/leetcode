// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
package main

type Node struct{ x, y int }

func (n Node) AdjacentTo(m Node) bool { return n.x == m.x || n.y == m.y }

type Set struct{ parent map[Node]Node }

func (this *Set) Union(n, m Node) {
	if pn, pm := this.Find(n), this.Find(m); pn != pm {
		this.parent[pm] = pn
	}
}

func (this *Set) Find(n Node) Node {
	if m, found := this.parent[n]; !found {
		this.parent[n] = n
	} else if n != m {
		this.parent[n] = this.Find(this.parent[n])
	}
	return this.parent[n]
}

func removeStones(stones [][]int) (parents int) {
	set, nodes := Set{make(map[Node]Node)}, make([]Node, len(stones))
	for i := range stones {
		nodes[i] = Node{stones[i][0], stones[i][1]}
	}
	for i := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].AdjacentTo(nodes[j]) {
				set.Union(nodes[i], nodes[j])
			}
		}
	}
	visited := make(map[Node]bool)
	for i := range nodes {
		if parent := set.Find(nodes[i]); !visited[parent] {
			visited[parent], parents = true, parents+1
		}
	}
	return len(nodes) - parents
}

func main() {}
