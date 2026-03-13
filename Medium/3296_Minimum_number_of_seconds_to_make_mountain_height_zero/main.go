// https://leetcode.com/problems/minimum-number-of-seconds-to-make-mountain-height-zero/
package main

type Item struct{ val, mul int64 }

type PriorityQueue struct{ items []Item }

func calc(item Item) int64 {
	return item.val * item.mul * (item.mul + 1) / 2
}

func (pq PriorityQueue) less(i, j int) bool {
	return calc(pq.items[i]) < calc(pq.items[j])
}

func NewPQ(items []Item) *PriorityQueue {
	pq := &PriorityQueue{append([]Item{{}}, items...)}
	for i := len(pq.items) / 2; i > 0; i-- {
		pq.heapify(i)
	}
	return pq
}

func (pq *PriorityQueue) heapify(i int) {
	lowest := i
	if l := i * 2; l < len(pq.items) && pq.less(l, lowest) {
		lowest = l
	}
	if r := i*2 + 1; r < len(pq.items) && pq.less(r, lowest) {
		lowest = r
	}
	if lowest != i {
		pq.items[i], pq.items[lowest] = pq.items[lowest], pq.items[i]
		pq.heapify(lowest)
	}
}

func (pq *PriorityQueue) Push(item Item) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.less(i, i/2); i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *PriorityQueue) Pop() Item {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.less(j+1, j) {
			j++
		}
		if pq.less(i, j) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) (res int64) {
	items := make([]Item, 0, len(workerTimes))
	for _, ts := range workerTimes {
		items = append(items, Item{int64(ts), 1})
	}
	for pq := NewPQ(items); mountainHeight > 0; mountainHeight-- {
		item := pq.Pop()
		res = max(res, calc(item))
		item.mul++
		pq.Push(item)
	}
	return res
}

func main() {}
