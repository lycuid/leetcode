// https://leetcode.com/problems/minimum-path-sum/
package main

type Tuple struct{ y, x, value int }
type PriorityQueue struct{ inner []Tuple }

func MakePQ() PriorityQueue                   { return PriorityQueue{make([]Tuple, 1)} }
func (this PriorityQueue) Empty() bool        { return len(this.inner) == 1 }
func (this PriorityQueue) Less(i, j int) bool { return this.inner[i].value < this.inner[j].value }

func (this *PriorityQueue) Push(item Tuple) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.Less(i, i/2); i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() Tuple {
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

func minPathSum(grid [][]int) int {
	visited, m, n := make([][]bool, len(grid)), len(grid), len(grid[0])
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	pq := MakePQ()
	for pq.Push(Tuple{0, 0, grid[0][0]}); !pq.Empty(); {
		if root := pq.Pop(); !visited[root.y][root.x] {
			visited[root.y][root.x], grid[root.y][root.x] = true, root.value
			if root.y == m-1 && root.x == n-1 {
				break
			}
			for _, d := range [][]int{{0, 1}, {1, 0}} {
				dy, dx := root.y+d[0], root.x+d[1]
				if dy < m && dx < n && !visited[dy][dx] {
					pq.Push(Tuple{dy, dx, root.value + grid[dy][dx]})
				}
			}
		}
	}
	return grid[m-1][n-1]
}

func main() {}
