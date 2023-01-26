// https://leetcode.com/problems/cheapest-flights-within-k-stops/
package main

type Tuple struct{ node, stops, weight int }
type PriorityQueue struct{ inner []Tuple }

func MakePQ() PriorityQueue                   { return PriorityQueue{make([]Tuple, 1)} }
func (this PriorityQueue) Empty() bool        { return len(this.inner) == 1 }
func (this PriorityQueue) Less(i, j int) bool { return this.inner[i].weight < this.inner[j].weight }
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

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	pq, adj, stops, tup, ret := MakePQ(), make([][]int, n), make([]int, n), Tuple{node: src}, -1
	for i := range adj {
		adj[i] = make([]int, n)
	}
	for _, f := range flights {
		adj[f[0]][f[1]] = f[2]
	}
	for i := range stops {
		stops[i] = k + 1
	}
	for pq.Push(tup); !pq.Empty(); {
		if tup = pq.Pop(); tup.node == dst {
			if ret == -1 || ret > tup.weight {
				ret = tup.weight
			}
		}
		if tup.stops < stops[tup.node] && tup.stops <= k {
			stops[tup.node] = tup.stops + 1
			for child, weight := range adj[tup.node] {
				if weight > 0 {
					pq.Push(Tuple{
						node:   child,
						stops:  tup.stops + 1,
						weight: tup.weight + weight,
					})
				}
			}
		}
	}
	return ret
}

func main() {}
