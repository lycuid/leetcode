// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
package main

type Coordinate struct {
	x, y int
}

type PriorityQueue struct {
	inner  [4]Coordinate
	cursor int
}

func (this *PriorityQueue) Push(item Coordinate, graph [][]int) {
	this.inner[this.cursor] = item
	for i := this.cursor; i > 0; i-- {
		f, s := this.inner[i], this.inner[i-1]
		if graph[f.x][f.y] > graph[s.x][s.y] {
			this.inner[i], this.inner[i-1] = this.inner[i-1], this.inner[i]
		}
	}
	this.cursor++
}

func (this *PriorityQueue) Pop() (item Coordinate) {
	this.cursor--
	item = this.inner[this.cursor]
	return item
}

func (this PriorityQueue) Empty() bool {
	return this.cursor == 0
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(x, y int, graph, weight [][]int, visited [][]bool) {
	pq := PriorityQueue{}
	visited[x][y] = true
	for _, d := range []int{-1, 1} {
		dx := x + d
		if dx >= 0 && dx < len(graph) && !visited[dx][y] {
			pq.Push(Coordinate{dx, y}, graph)
		}
	}
	for _, d := range []int{-1, 1} {
		dy := y + d
		if dy >= 0 && dy < len(graph[0]) && !visited[x][dy] {
			pq.Push(Coordinate{x, dy}, graph)
		}
	}
	for !pq.Empty() {
		c := pq.Pop()
		if graph[c.x][c.y] == graph[x][y] {
			continue
		}
		if graph[c.x][c.y] > graph[x][y] {
			weight[c.x][c.y] = Max(weight[c.x][c.y], weight[x][y]+1)
			continue
		}
		dfs(c.x, c.y, graph, weight, visited)
		weight[x][y] = Max(weight[x][y], weight[c.x][c.y]+1)
	}
}

func longestIncreasingPath(graph [][]int) (count int) {
	weight := make([][]int, len(graph))
	for i := range weight {
		weight[i] = make([]int, len(graph[0]))
		for j := range weight[i] {
			weight[i][j] = 1
		}
	}
	visited := make([][]bool, len(graph))
	for i := range visited {
		visited[i] = make([]bool, len(graph[0]))
	}
	for x := 0; x < len(graph); x++ {
		for y := 0; y < len(graph[0]); y++ {
			if !visited[x][y] {
				dfs(x, y, graph, weight, visited)
			}
		}
	}
	for i := range weight {
		for _, w := range weight[i] {
			count = Max(count, w)
		}
	}
	return count
}

func main() {}
