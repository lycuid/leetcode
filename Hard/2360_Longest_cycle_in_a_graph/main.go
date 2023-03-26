// https://leetcode.com/problems/longest-cycle-in-a-graph/
package main

type Stack struct {
	inner  []int
	cursor int
	has    []bool
}

func MakeStack(n int) Stack {
	return Stack{inner: make([]int, n), has: make([]bool, n)}
}

func (this Stack) Empty() bool { return this.cursor == 0 }

func (this *Stack) Push(item int) {
	this.inner[this.cursor], this.cursor = item, this.cursor+1
	this.has[item] = true
}

func (this *Stack) Pop() (item int) {
	item, this.cursor = this.inner[this.cursor-1], this.cursor-1
	this.has[item] = false
	return item
}

func longestCycle(edges []int) int {
	done, ret, n := make([]bool, len(edges)), -1, len(edges)
	for stack, root := MakeStack(n), 0; root < n; root++ {
		node := root
		for ; node != -1 && !done[node] && !stack.has[node]; node = edges[node] {
			stack.Push(node)
		}
		for size := 0; !stack.Empty(); {
			item := stack.Pop()
			if size++; item == node && size > ret {
				ret = size
			}
			done[item] = true
		}
	}
	return ret
}

func main() {}
