// https://leetcode.com/problems/merge-intervals/
package main

import "sort"

type Stack struct {
	inner  [][]int
	cursor int
}

func MakeStack(n int) Stack {
	return Stack{inner: make([][]int, n)}
}

func (this *Stack) Push(item []int) {
	this.inner[this.cursor] = item
	this.cursor++
}

func (this *Stack) Top() *[]int {
	if this.cursor == 0 {
		return nil
	}
	return &this.inner[this.cursor-1]
}

func (this *Stack) Get() [][]int {
	return this.inner[:this.cursor]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	stack := MakeStack(len(intervals))
	for _, interval := range intervals {
		top := stack.Top()
		if top != nil && (*top)[1] >= interval[0] {
			(*top)[1] = Max((*top)[1], interval[1])
		} else {
			stack.Push(interval)
		}
	}
	return stack.Get()
}

func main() {}
