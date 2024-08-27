// https://leetcode.com/problems/path-with-maximum-probability/
package main

type Tuple struct {
	node   int
	weight float64
}

type PriorityQueue struct{ items []Tuple }

func MakePQ() PriorityQueue {
	return PriorityQueue{make([]Tuple, 1)}
}

func (pq PriorityQueue) Empty() bool        { return len(pq.items) == 1 }
func (pq PriorityQueue) Less(i, j int) bool { return pq.items[i].weight > pq.items[j].weight }

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

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	weight := make([]map[int]float64, n)
	visited := make([]bool, n)
	for i, edge := range edges {
		if weight[edge[0]] == nil {
			weight[edge[0]] = make(map[int]float64)
		}
		weight[edge[0]][edge[1]] = succProb[i]
		if weight[edge[1]] == nil {
			weight[edge[1]] = make(map[int]float64)
		}
		weight[edge[1]][edge[0]] = succProb[i]
	}

	pq := MakePQ()
	pq.Push(Tuple{start_node, 1})
	for !pq.Empty() {
		parent := pq.Pop()
		if visited[parent.node] = true; parent.node == end_node {
			return parent.weight
		}
		for child, weight := range weight[parent.node] {
			if !visited[child] {
				pq.Push(Tuple{child, parent.weight * weight})
			}
		}
	}
	return 0
}

func main() {}
