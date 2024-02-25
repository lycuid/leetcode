// https://leetcode.com/problems/greatest-common-divisor-traversal/
package main

type Graph struct{ parent []int }

func MakeGraph(n int) *Graph {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &Graph{parent}
}

func (this *Graph) Union(u, v int) {
	if pu, pv := this.Find(u), this.Find(v); pu != pv {
		this.parent[pv] = pu
	}
}

func (this *Graph) Find(u int) int {
	if this.parent[u] != u {
		this.parent[u] = this.Find(this.parent[u])
	}
	return this.parent[u]
}

func Factors(num int) (factors []int) {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			for num%i == 0 {
				num /= i
			}
		}
	}
	if num > 1 {
		factors = append(factors, num)
	}
	return factors
}

func canTraverseAllPairs(nums []int) bool {
	groups := make(map[int][]int)
	for i := range nums {
		factors := Factors(nums[i])
		for _, factor := range factors {
			groups[factor] = append(groups[factor], i)
		}
	}

	graph := MakeGraph(len(nums))
	for _, member := range groups {
		for i := 1; i < len(member); i++ {
			graph.Union(member[i-1], member[i])
		}
	}

	root := graph.Find(0)
	for i := range nums {
		if graph.Find(i) != root {
			return false
		}
	}
	return true
}

func main() {}
