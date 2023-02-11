// https://leetcode.com/problems/shortest-path-with-alternating-colors/
package main

import "math"

type Point struct{ path, node, weight int }
type PriorityQueue struct{ inner []Point }

func MakePQ() PriorityQueue            { return PriorityQueue{make([]Point, 1)} }
func (this PriorityQueue) Empty() bool { return len(this.inner) == 1 }

func (this *PriorityQueue) Push(item Point) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i].weight < this.inner[i/2].weight; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() Point {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.inner[j+1].weight < this.inner[j].weight {
			j++
		}
		if this.inner[i].weight < this.inner[j].weight {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	var (
		Red       = 0
		Blue      = 1
		distances [2][]int
		adj       [2]map[int][]int
	)
	distances[Red], distances[Blue] = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		distances[Red][i], distances[Blue][i] = math.MaxInt, math.MaxInt
	}
	adj[Red], adj[Blue] = make(map[int][]int), make(map[int][]int)
	for _, e := range redEdges {
		adj[Red][e[0]] = append(adj[Red][e[0]], e[1])
	}
	for _, e := range blueEdges {
		adj[Blue][e[0]] = append(adj[Blue][e[0]], e[1])
	}
	pq := MakePQ()
	pq.Push(Point{Red, 0, 0})
	pq.Push(Point{Blue, 0, 0})
	for !pq.Empty() {
		if r := pq.Pop(); r.weight < distances[r.path][r.node] {
			distances[r.path][r.node] = r.weight
			for _, node := range adj[1-r.path][r.node] {
				pq.Push(Point{1 - r.path, node, r.weight + 1})
			}
		}
	}
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		if dist[i] = distances[Red][i]; dist[i] > distances[Blue][i] {
			dist[i] = distances[Blue][i]
		}
		if dist[i] == math.MaxInt {
			dist[i] = -1
		}
	}
	return dist
}

func main() {}
