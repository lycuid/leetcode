// https://leetcode.com/problems/similar-string-groups/
package main

type Graph struct{ parent []int }

func MakeGraph(n int) Graph {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return Graph{parent}
}

func (this *Graph) Find(x int) int {
	if x != this.parent[x] {
		this.parent[x] = this.Find(this.parent[x])
	}
	return this.parent[x]
}

func (this *Graph) Union(x, y int) {
	if px, py := this.Find(x), this.Find(y); px != py {
		this.parent[py] = px
	}
}

func Similar(s1, s2 string) bool {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count <= 2
}

func numSimilarGroups(strs []string) (ret int) {
	graph, n := MakeGraph(len(strs)), len(strs)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if Similar(strs[i], strs[j]) {
				graph.Union(i, j)
			}
		}
	}
	seen := make([]bool, n)
	for i := 0; i < n; i++ {
		if j := graph.Find(i); !seen[j] {
			ret, seen[j] = ret+1, true
		}
	}
	return ret
}

func main() {}
