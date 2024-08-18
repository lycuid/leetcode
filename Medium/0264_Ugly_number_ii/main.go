// https://leetcode.com/problems/ugly-number-ii/
package main

type PriorityQueue struct{ items []int }

func MakePQ() PriorityQueue          { return PriorityQueue{make([]int, 1)} }
func (pq PriorityQueue) Empty() bool { return len(pq.items) == 1 }
func (pq PriorityQueue) Top() int    { return pq.items[1] }

func (pq *PriorityQueue) Push(item int) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.items[i] < pq.items[i/2]; i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *PriorityQueue) Pop() int {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.items[j+1] < pq.items[j] {
			j++
		}
		if pq.items[i] < pq.items[j] {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

func nthUglyNumber(n int) int {
	pq := MakePQ()
	for pq.Push(1); n > 1; n-- {
		item := pq.Pop()
		for !pq.Empty() && pq.Top() == item {
			pq.Pop()
		}
		pq.Push(item * 2)
		pq.Push(item * 3)
		pq.Push(item * 5)
	}
	return pq.Pop()
}

func main() {}
