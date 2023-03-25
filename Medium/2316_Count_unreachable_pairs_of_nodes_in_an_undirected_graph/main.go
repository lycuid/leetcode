// https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
package main

type Graph struct{ inner []int }

func MakeGraph(n int) Graph {
	inner := make([]int, n)
	for i := range inner {
		inner[i] = i
	}
	return Graph{inner}
}

func (this *Graph) Union(x, y int) {
	if px, py := this.Find(x), this.Find(y); px != py {
		this.inner[py] = px
	}
}

func (this *Graph) Find(x int) int {
	if x != this.inner[x] {
		this.inner[x] = this.Find(this.inner[x])
	}
	return this.inner[x]
}

func countPairs(n int, edges [][]int) (ret int64) {
	graph := MakeGraph(n)
	for _, e := range edges {
		graph.Union(e[0], e[1])
	}
	for cache, sum, i := make([]int, n), 0, 0; i < n; i++ {
		p := graph.Find(i)
		ret, cache[p], sum = ret+int64(sum-cache[p]), cache[p]+1, sum+1
	}
	return ret
}

func main() {}
