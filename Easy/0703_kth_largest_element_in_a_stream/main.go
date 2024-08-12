// https://leetcode.com/problems/kth-largest-element-in-a-stream/
package main

type PriorityQueue struct{ items []int }

func MakePQ(nums []int) (pq PriorityQueue) {
	nums = append([]int{0}, nums...)
	for i := len(nums) / 2; i > 0; i-- {
		Heapify(nums, i)
	}
	return PriorityQueue{nums}
}

func Heapify(nums []int, i int) {
	if j := i * 2; j < len(nums) {
		if j+1 < len(nums) && nums[j+1] < nums[j] {
			j++
		}
		if nums[j] < nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
			Heapify(nums, j)
		}
	}
}

func (pq PriorityQueue) Len() int { return len(pq.items) - 1 }

func (pq PriorityQueue) Top() int { return pq.items[1] }

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

type KthLargest struct {
	k  int
	pq PriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{k: k, pq: MakePQ(nums)}
}

func (this *KthLargest) Add(val int) int {
	this.pq.Push(val)
	for this.pq.Len() > this.k {
		this.pq.Pop()
	}
	return this.pq.Top()
}

func main() {}
