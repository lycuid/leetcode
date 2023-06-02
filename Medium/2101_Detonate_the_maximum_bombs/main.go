// https://leetcode.com/problems/detonate-the-maximum-bombs/
package main

type Solution struct {
	nodes   [][]int
	visited []bool
}

func MakeSolution(bombs [][]int) *Solution {
	return &Solution{bombs, make([]bool, len(bombs))}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (this Solution) Reachable(from, to int) bool {
	x := Abs(this.nodes[from][0] - this.nodes[to][0])
	y := Abs(this.nodes[from][1] - this.nodes[to][1])
	r := this.nodes[from][2]
	return r*r >= (x*x)+(y*y)
}

func (this *Solution) Solve() (count int) {
	for i := range this.nodes {
		this.Dfs(i)
		n := 0
		for j := range this.visited {
			if this.visited[j] {
				n++
			}
			this.visited[j] = false
		}
		if n > count {
			count = n
		}
	}
	return count
}

func (this *Solution) Dfs(root int) {
	this.visited[root] = true
	for i := range this.nodes {
		if !this.visited[i] && this.Reachable(root, i) {
			this.Dfs(i)
		}
	}
}

func maximumDetonation(bombs [][]int) int {
	solution := MakeSolution(bombs)
	return solution.Solve()
}

func main() {}
