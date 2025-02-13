// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/
package main

type PriorityQueue struct{ items []int }

func MakePQ(items []int) PriorityQueue {
	items = append([]int{0}, items...)
	for i := len(items)/2 + 1; i > 0; i-- {
		heapify(items, i)
	}
	return PriorityQueue{items}
}

func heapify(items []int, index int) {
	pIndex := index
	if i := index * 2; i < len(items) && items[i] < items[pIndex] {
		pIndex = i
	}
	if i := index*2 + 1; i < len(items) && items[i] < items[pIndex] {
		pIndex = i
	}
	if pIndex != index {
		items[index], items[pIndex] = items[pIndex], items[index]
		heapify(items, pIndex)
	}
}

func (pq *PriorityQueue) Top() *int {
	if len(pq.items) == 1 {
		return nil
	}
	return &pq.items[1]
}

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

func minOperations(nums []int, k int) (count int) {
	pq := MakePQ(nums)
	for ; pq.Top() != nil && *pq.Top() < k; count++ {
		fst, snd := pq.Pop(), pq.Pop()
		pq.Push(min(fst, snd)*2 + max(fst, snd))
	}
	return count
}

func main() {}
