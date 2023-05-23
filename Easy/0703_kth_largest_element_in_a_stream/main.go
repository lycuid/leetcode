// https://leetcode.com/problems/kth-largest-element-in-a-stream/
package main

type PriorityQueue struct{ inner []int }

func MakePQ() PriorityQueue         { return PriorityQueue{make([]int, 1)} }
func (this PriorityQueue) Len() int { return len(this.inner) - 1 }
func (this PriorityQueue) Top() int { return this.inner[1] }

func (this *PriorityQueue) Push(item int) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i] < this.inner[i/2]; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() int {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.inner[j+1] < this.inner[j] {
			j++
		}
		if this.inner[i] < this.inner[j] {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

type KthLargest struct {
	k  int
	pq PriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
	pq := MakePQ()
	for _, num := range nums {
		pq.Push(num)
	}
	return KthLargest{k, pq}
}

func (this *KthLargest) Add(val int) int {
	for this.pq.Push(val); this.pq.Len() > this.k; {
		this.pq.Pop()
	}
	return this.pq.Top()
}

func main() {}
