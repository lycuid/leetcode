// https://leetcode.com/problems/lexicographically-smallest-equivalent-string/description/
package main

type Graph struct{ parent [26]int }

func NewGraph() *Graph {
	var parent [26]int
	for i := range parent {
		parent[i] = i
	}
	return &Graph{parent}
}

func (g *Graph) Union(x, y int) {
	if px, py := g.Find(x), g.Find(y); px != py {
		if px < py {
			g.parent[py] = px
		} else {
			g.parent[px] = py
		}
	}
}

func (g *Graph) Find(x int) int {
	if x != g.parent[x] {
		g.parent[x] = g.Find(g.parent[x])
	}
	return g.parent[x]
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	var (
		graph = NewGraph()
		res   = make([]byte, len(baseStr))
	)
	for i := 0; i < len(s1); i++ {
		graph.Union(int(s1[i]-'a'), int(s2[i]-'a'))
	}
	for i := range res {
		res[i] = byte(graph.Find(int(baseStr[i]-'a'))) + 'a'
	}
	return string(res)
}

func main() {}
