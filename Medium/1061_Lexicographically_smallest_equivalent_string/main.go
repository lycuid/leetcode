// https://leetcode.com/problems/lexicographically-smallest-equivalent-string/
package main

type Graph struct{ parent [26]int }

func MakeGraph() (graph Graph) {
	for i := range graph.parent {
		graph.parent[i] = i
	}
	return graph
}

func (this *Graph) Union(x, y int) {
	if px, py := this.Find(this.parent[x]), this.Find(this.parent[y]); px != py {
		if px < py {
			this.parent[py] = px
		} else {
			this.parent[px] = py
		}
	}
}

func (this *Graph) Find(x int) int {
	if x != this.parent[x] {
		this.parent[x] = this.Find(this.parent[x])
	}
	return this.parent[x]
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	graph, ret := MakeGraph(), make([]byte, len(baseStr))
	for i := range s1 {
		graph.Union(int(s1[i]-'a'), int(s2[i]-'a'))
	}
	for i, ch := range baseStr {
		ret[i] = byte(graph.Find(int(ch-'a')) + 'a')
	}
	return string(ret)
}

func main() {}
