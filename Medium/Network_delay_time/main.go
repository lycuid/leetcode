// https://leetcode.com/problems/network-delay-time/
package main

import "math"

type Queue struct {
	inner   []int
	visited []bool
}

func MakeQueue(n int) Queue {
	return Queue{visited: make([]bool, n)}
}

func (this *Queue) Put(item int) {
	if !this.visited[item] {
		this.inner = append(this.inner, item)
		this.visited[item] = true
	}
}

func (this *Queue) Pop(path []int) (item int) {
	index := 0
	for i := 0; i < len(this.inner); i++ {
		if path[this.inner[i]] < path[this.inner[index]] {
			index = i
		}
	}
	this.inner[0], this.inner[index] = this.inner[index], this.inner[0]
	item = this.inner[0]
	this.inner = this.inner[1:]
	return item
}

func (this Queue) Empty() bool {
	return len(this.inner) == 0
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func networkDelayTime(times [][]int, n int, k int) int {
	src := k - 1
	edge, path := make([][]int, n), make([]int, n)
	for i := range edge {
		edge[i] = make([]int, n)
		for j := range edge[i] {
			edge[i][j] = math.MaxInt
		}
	}
	for i := 0; i < n; i++ {
		path[i] = math.MaxInt
	}
	for _, time := range times {
		u, v, w := time[0]-1, time[1]-1, time[2]
		edge[u][v] = w
	}
	edge[src][src], path[src] = 0, 0

	q := MakeQueue(n)
	q.Put(src)
	for !q.Empty() {
		root := q.Pop(path)
		for i, weight := range edge[root] {
			if weight == math.MaxInt || i == root {
				continue
			}
			q.Put(i)
			path[i] = Min(path[i], edge[root][i]+path[root])
		}
	}
	for i := 1; i < n; i++ {
		path[i] = Max(path[i], path[i-1])
	}
	if path[n-1] == math.MaxInt {
		return -1
	}
	return path[n-1]
}

func main() {}
