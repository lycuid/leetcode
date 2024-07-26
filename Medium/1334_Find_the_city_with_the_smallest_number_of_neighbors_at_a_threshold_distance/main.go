// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
package main

import "math"

type Tuple struct{ node, weight int }
type PriorityQueue struct{ items []Tuple }

func MakePQ() PriorityQueue                 { return PriorityQueue{make([]Tuple, 1)} }
func (pq PriorityQueue) Empty() bool        { return len(pq.items) == 1 }
func (pq PriorityQueue) Less(i, j int) bool { return pq.items[i].weight < pq.items[j].weight }

func (pq *PriorityQueue) Push(item Tuple) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.Less(i, i/2); i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *PriorityQueue) Pop() Tuple {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.Less(j+1, j) {
			j++
		}
		if pq.Less(i, j) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

type Solution struct {
	dist    []map[int]int
	weight  []int
	visited []bool
}

func MakeSolution(n int, edges [][]int) Solution {
	dist := make([]map[int]int, n)
	for i := range dist {
		dist[i] = make(map[int]int)
		dist[i][i] = 0
	}
	for _, edge := range edges {
		dist[edge[0]][edge[1]] = edge[2]
		dist[edge[1]][edge[0]] = edge[2]
	}
	return Solution{dist: dist, weight: make([]int, n), visited: make([]bool, n)}
}

func (this *Solution) GetCityCount(node, fuel int) (count int) {
	pq := MakePQ()
	for i := 0; i < len(this.weight); i++ {
		this.weight[i], this.visited[i] = math.MaxInt, false
	}
	this.weight[node] = 0
	pq.Push(Tuple{node, 0})
	for !pq.Empty() {
		parent := pq.Pop()
		if this.visited[parent.node] {
			continue
		}
		this.visited[parent.node] = true
		for child, dist := range this.dist[parent.node] {
			if dist > fuel || this.visited[child] {
				continue
			}
			this.weight[child] = min(this.weight[child], this.weight[parent.node]+dist)
			if this.weight[child] <= fuel {
				pq.Push(Tuple{child, this.weight[child]})
			}
		}
	}
	for _, dist := range this.weight {
		if dist <= fuel {
			count++
		}
	}
	return count - 1
}

func findTheCity(n int, edges [][]int, distanceThreshold int) (city int) {
	solution := MakeSolution(n, edges)
	for i, count := 0, math.MaxInt; i < n; i++ {
		if c := solution.GetCityCount(i, distanceThreshold); c <= count {
			count, city = c, i
		}
	}
	return city
}

func main() {}
