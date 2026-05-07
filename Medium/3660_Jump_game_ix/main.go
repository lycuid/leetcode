// https://leetcode.com/problems/jump-game-ix/
package main

type Graph struct{ nums, parent []int }

func makeGraph(nums []int) Graph {
	parent := make([]int, len(nums))
	for i := range parent {
		parent[i] = i
	}
	return Graph{nums: nums, parent: parent}
}

func (g *Graph) union(x, y int) {
	if px, py := g.find(x), g.find(y); px != py {
		if g.nums[px] > g.nums[py] {
			g.parent[py] = px
		} else {
			g.parent[px] = py
		}
	}
}

func (g *Graph) find(x int) int {
	if x != g.parent[x] {
		g.parent[x] = g.find(g.parent[x])
	}
	return g.parent[x]
}

func maxValue(nums []int) []int {
	graph := makeGraph(nums)
	prefix, suffix := make([]int, len(nums)), len(nums)-1
	for i := 1; i < len(nums); i++ {
		if prefix[i] = prefix[i-1]; nums[i] < nums[prefix[i]] {
			graph.union(i, prefix[i])
		} else {
			prefix[i] = i
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[suffix] {
			graph.union(i, suffix)
		} else {
			if nums[prefix[i]] > nums[suffix] {
				graph.union(suffix, prefix[i])
			}
			suffix = i
		}
	}
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = nums[graph.find(i)]
	}
	return res
}

func main() {}
