// https://leetcode.com/problems/as-far-from-land-as-possible/
package main

type Point struct{ y, x, value int }
type PriorityQueue struct{ inner []Point }

func MakePQ() PriorityQueue { return PriorityQueue{make([]Point, 1)} }

func (this PriorityQueue) Empty() bool { return len(this.inner) == 1 }
func (this *PriorityQueue) Push(item Point) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i].value < this.inner[i/2].value; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() Point {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.inner[j+1].value < this.inner[j].value {
			j++
		}
		if this.inner[i].value < this.inner[j].value {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func maxDistance(grid [][]int) (ret int) {
	pq, n := MakePQ(), len(grid)
	MAX := n*n + 1
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				grid[y][x] = MAX
			}
			if grid[y][x] == 1 {
				grid[y][x] = -1
			}
		}
	}
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == -1 {
				for _, d := range [][]int{{y, x + 1}, {y, x - 1}, {y + 1, x}, {y - 1, x}} {
					dy, dx := d[0], d[1]
					if dy >= 0 && dy < n && dx >= 0 && dx < n && grid[dy][dx] != -1 {
						grid[dy][dx] = 1
						pq.Push(Point{dy, dx, grid[dy][dx]})
					}
				}
			}
		}
	}
	for ret = -1; !pq.Empty(); {
		p := pq.Pop()
		for _, d := range [][]int{{p.y, p.x + 1}, {p.y, p.x - 1}, {p.y + 1, p.x}, {p.y - 1, p.x}} {
			if dy, dx := d[0], d[1]; dy >= 0 && dy < n && dx >= 0 && dx < n {
				if grid[p.y][p.x]+1 < grid[dy][dx] {
					grid[dy][dx] = grid[p.y][p.x] + 1
					pq.Push(Point{dy, dx, grid[dy][dx]})
				}
			}
		}
	}
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] < MAX && grid[y][x] > ret {
				ret = grid[y][x]
			}
		}
	}
	return ret
}

func main() {}
