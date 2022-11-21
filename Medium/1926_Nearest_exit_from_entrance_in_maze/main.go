// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
package main

const PROCESSED = '+'
const INF = 1e7 + 9

type Point struct{ y, x int }
type PriorityQueue struct {
	inner  []Point
	weight [][]int
}

func MakePQ(r, c int) PriorityQueue {
	weight := make([][]int, r)
	for i := range weight {
		weight[i] = make([]int, c)
		for j := range weight[i] {
			weight[i][j] = INF
		}
	}
	return PriorityQueue{make([]Point, 1), weight}
}

func (this PriorityQueue) Empty() bool { return len(this.inner) == 1 }
func (this PriorityQueue) Less(i, j int) bool {
	c1, c2 := this.inner[i], this.inner[j]
	return this.weight[c1.y][c1.x] < this.weight[c2.y][c2.x]
}

func (this *PriorityQueue) Push(p Point, w int) {
	this.inner, this.weight[p.y][p.x] = append(this.inner, p), w
	for i := len(this.inner) - 1; i > 1 && this.Less(i, i/2); i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() Point {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.Less(j+1, j) {
			j++
		}
		if this.Less(i, j) {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func Children(p Point, grid [][]byte) (children []Point) {
	for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		dx, dy, r, c := p.x+d[0], p.y+d[1], len(grid), len(grid[0])
		if dx >= 0 && dx < c && dy >= 0 && dy < r && grid[dy][dx] != PROCESSED {
			children = append(children, Point{dy, dx})
		}
	}
	return children
}

func nearestExit(maze [][]byte, entrance []int) int {
	n, m := len(maze), len(maze[0])
	pq, root := MakePQ(n, m), Point{entrance[0], entrance[1]}
	for pq.Push(root, 0); !pq.Empty(); {
		r := pq.Pop()
		if r != root && (r.x == 0 || r.x == m-1 || r.y == 0 || r.y == n-1) {
			return pq.weight[r.y][r.x]
		}
		maze[r.y][r.x] = PROCESSED
		for _, c := range Children(r, maze) {
			if pq.weight[c.y][c.x] > pq.weight[r.y][r.x]+1 {
				pq.Push(c, pq.weight[r.y][r.x]+1)
			}
		}
	}
	return -1
}

func main() {}
