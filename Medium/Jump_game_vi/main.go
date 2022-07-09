// https://leetcode.com/problems/jump-game-vi/
package main

type Tuple struct {
	index, value int
}

type PriorityQueue struct {
	inner []Tuple
}

func MakePQ() PriorityQueue {
	return PriorityQueue{make([]Tuple, 1)}
}

func (this *PriorityQueue) Top() Tuple { return this.inner[1] }

func (this *PriorityQueue) Push(item Tuple) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i].value > this.inner[i/2].value; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() Tuple {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1] = this.inner[n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if this.inner[j].value < this.inner[j+1].value {
			j++
		}
		if this.inner[j].value < this.inner[i].value {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	this.inner = this.inner[:n]
	return item
}
func (this PriorityQueue) Empty() bool {
	return len(this.inner) == 1
}

func maxResult(nums []int, k int) int {
	n, pq := len(nums), MakePQ()
	cache := make([]int, len(nums))
	for i := n - 1; i >= 0; i-- {
		for !pq.Empty() && pq.Top().index > i+k {
			pq.Pop()
		}
		cache[i] = nums[i]
		if !pq.Empty() {
			cache[i] += pq.Top().value
		}
		pq.Push(Tuple{i, cache[i]})
	}
	return cache[0]
}

func main() {}
