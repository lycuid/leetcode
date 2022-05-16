// https://leetcode.com/problems/shortest-path-in-binary-matrix/
package main

import "math"

type Coordinate struct {
	x, y int
}

type Queue struct {
	inner      []Coordinate
	start, end int
	visited    [][]bool
}

func MakeQueue(n int) Queue {
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	return Queue{inner: make([]Coordinate, n*n), visited: visited}
}

func (this *Queue) Put(item Coordinate) {
	if !this.visited[item.x][item.y] {
		this.inner[this.end] = item
		this.end = (this.end + 1) % len(this.inner)
		this.visited[item.x][item.y] = true
	}
}

func (this *Queue) Pop() (item Coordinate) {
	item = this.inner[this.start]
	this.start = (this.start + 1) % len(this.inner)
	return item
}

func (this *Queue) Empty() bool {
	return this.start == this.end
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	L := len(grid)
	src, weights := Coordinate{0, 0}, make([][]int, L)
	for i := range weights {
		weights[i] = make([]int, L)
		for j := range weights[i] {
			weights[i][j] = math.MaxInt
		}
	}
	weights[src.x][src.y] = 1

	q := MakeQueue(L)
	q.Put(src)
	for !q.Empty() {
		node := q.Pop()
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				cx, cy := node.x+x, node.y+y
				if cx >= 0 && cx < L && cy >= 0 && cy < L && grid[cx][cy] != 1 {
					q.Put(Coordinate{cx, cy})
					weights[cx][cy] = Min(weights[cx][cy], weights[node.x][node.y]+1)
				}
			}
		}
	}

	weight := weights[len(grid)-1][len(grid)-1]
	if weight == math.MaxInt {
		return -1
	}
	return weight
}

func main() {}
