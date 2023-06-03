// https://leetcode.com/problems/time-needed-to-inform-all-employees/
package main

type Solution struct {
	tree map[int][]int
	time []int
}

func MakeSolution(manager []int, time []int) *Solution {
	tree := make(map[int][]int)
	for child, parent := range manager {
		if parent != -1 {
			tree[parent] = append(tree[parent], child)
		}
	}
	return &Solution{tree, time}
}

func (this Solution) Solve(parent int) (count int) {
	for _, child := range this.tree[parent] {
		if n := this.Solve(child); n > count {
			count = n
		}
	}
	return count + this.time[parent]
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	solution := MakeSolution(manager, informTime)
	return solution.Solve(headID)
}

func main() {}
