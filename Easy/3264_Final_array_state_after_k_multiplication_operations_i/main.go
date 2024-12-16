// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/
package main

type PriorityQueue struct {
	items []int
	less  func(int, int) bool
}

func MakePQ(items []int, less func(int, int) bool) (pq PriorityQueue) {
	pq.less = less
	pq.items = append([]int{0}, items...)
	for i := len(pq.items) / 2; i > 0; i-- {
		pq.heapify(i)
	}
	return pq
}

func (pq *PriorityQueue) heapify(curr int) {
	index, n := curr, len(pq.items)
	if i := curr * 2; i < n && pq.less(pq.items[i], pq.items[index]) {
		index = i
	}
	if i := curr*2 + 1; i < n && pq.less(pq.items[i], pq.items[index]) {
		index = i
	}
	if index != curr {
		pq.items[index], pq.items[curr] = pq.items[curr], pq.items[index]
		pq.heapify(index)
	}
}

func (pq *PriorityQueue) Push(item int) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.less(pq.items[i], pq.items[i/2]); i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *PriorityQueue) Pop() int {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.less(pq.items[j+1], pq.items[j]) {
			j++
		}
		if pq.less(pq.items[i], pq.items[j]) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

func getFinalState(nums []int, k int, multiplier int) []int {
	items := make([]int, len(nums))
	for i := range items {
		items[i] = i
	}
	pq := MakePQ(items, func(i, j int) bool {
		if nums[i] == nums[j] {
			return i < j
		}
		return nums[i] < nums[j]
	})
	for ; k > 0; k-- {
		i := pq.Pop()
		nums[i] *= multiplier
		pq.Push(i)
	}
	return nums
}

func main() {}
