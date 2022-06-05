// https://leetcode.com/problems/minimum-path-sum/
package main

import "math"

type Coordinate struct {
	x, y int
}

type PriorityQueue struct {
	inner  [200*200 + 1]Coordinate
	cursor int
}

func (this *PriorityQueue) Push(item Coordinate, weight [][]int) {
	this.inner[this.cursor] = item
	for i := this.cursor; i > 0; i-- {
		c1, c2 := this.inner[i], this.inner[i-1]
		if weight[c1.x][c1.y] < weight[c2.x][c2.y] {
			break
		}
		this.inner[i], this.inner[i-1] = this.inner[i-1], this.inner[i]
	}
	this.cursor++
}

func (this *PriorityQueue) Pop() Coordinate {
	this.cursor--
	return this.inner[this.cursor]
}

func (this *PriorityQueue) Empty() bool {
	return this.cursor == 0
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	weight := make([][]int, m)
	for i := range weight {
		weight[i] = make([]int, n)
		for j := range weight[i] {
			weight[i][j] = math.MaxInt
		}
	}

	pq, src := PriorityQueue{}, Coordinate{0, 0}
	weight[0][0] = grid[0][0]
	pq.Push(src, weight)
	for !pq.Empty() {
		p := pq.Pop()
		if p.x+1 < m {
			c := Coordinate{p.x + 1, p.y}
			visited := weight[c.x][c.y] != math.MaxInt
			weight[c.x][c.y] = Min(weight[c.x][c.y], weight[p.x][p.y]+grid[c.x][c.y])
			if !visited {
				pq.Push(c, weight)
			}
		}
		if p.y+1 < n {
			c := Coordinate{p.x, p.y + 1}
			visited := weight[c.x][c.y] != math.MaxInt
			weight[c.x][c.y] = Min(weight[c.x][c.y], weight[p.x][p.y]+grid[c.x][c.y])
			if !visited {
				pq.Push(c, weight)
			}
		}
	}
	return weight[m-1][n-1]
}

func main() {}
