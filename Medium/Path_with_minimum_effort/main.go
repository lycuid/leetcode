// https://leetcode.com/problems/path-with-minimum-effort/
package main

import "math"

type Coordinate struct {
	x, y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

type PriorityQueue struct {
	queue    []Coordinate
	priority []int
}

func (pq *PriorityQueue) Put(c Coordinate, p int) {
	pq.queue = append(pq.queue, c)
	pq.priority = append(pq.priority, p)
}

func (pq *PriorityQueue) Pop() (Coordinate, int) {
	top := 0
	for i, priority := range pq.priority {
		if priority < pq.priority[top] {
			top = i
		}
	}
	c, p := pq.queue[top], pq.priority[top]
	pq.queue[0], pq.queue[top] = pq.queue[top], pq.queue[0]
	pq.priority[0], pq.priority[top] = pq.priority[top], pq.priority[0]
	pq.queue, pq.priority = pq.queue[1:], pq.priority[1:]
	return c, p
}

func (pq PriorityQueue) Empty() bool {
	return len(pq.queue) == 0
}

func minimumEffortPath(heights [][]int) int {
	rows, cols := len(heights), len(heights[0])
	visited, weights := make(map[Coordinate]bool), make([][]int, rows)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			visited[Coordinate{row, col}] = false
		}
	}
	for i := range weights {
		weights[i] = make([]int, cols)
		for j := range weights[i] {
			weights[i][j] = math.MaxInt
		}
	}
	weights[0][0] = 0

	pq, delta := PriorityQueue{}, []Coordinate{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	pq.Put(Coordinate{0, 0}, 0)
	for !pq.Empty() {
		coord, w := pq.Pop()
		if visited[coord] {
			continue
		}
		visited[coord] = true
		x, y := coord.x, coord.y
		for _, d := range delta {
			dx, dy := x+d.x, y+d.y
			if dx >= 0 && dx < rows && dy >= 0 && dy < cols {
				weight := Max(w, Abs(heights[x][y]-heights[dx][dy]))
				weights[dx][dy] = Min(weight, weights[dx][dy])
				if !visited[Coordinate{dx, dy}] {
					pq.Put(Coordinate{dx, dy}, weight)
				}
			}
		}
	}
	return weights[rows-1][cols-1]
}

func main() {}
