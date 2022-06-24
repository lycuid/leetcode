// https://leetcode.com/problems/construct-target-array-with-multiple-Sums/
package main

type PriorityQueue struct {
	inner []int
	Sum   int
}

func MakePQ() PriorityQueue         { return PriorityQueue{inner: make([]int, 1)} }
func (this PriorityQueue) Top() int { return this.inner[1] }

func (this *PriorityQueue) Push(item int) {
	this.Sum += item
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i/2] < this.inner[i]; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() int {
	item, n := this.inner[1], len(this.inner)
	this.inner[1], this.inner = this.inner[n-1], this.inner[:n-1]
	for i, j := 1, 2; j < n-1; i, j = j, j*2 {
		if j+1 < n-1 && this.inner[j+1] > this.inner[j] {
			j++
		}
		if this.inner[i] > this.inner[j] {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	this.Sum -= item
	return item
}

func isPossible(target []int) bool {
	if len(target) == 1 && target[0] > 1 {
		return false
	}
	pq := MakePQ()
	for _, item := range target {
		pq.Push(item)
	}
	for item := pq.Pop(); item > 1; item = pq.Pop() {
		if item <= pq.Sum {
			return false
		}
		if pq.Sum > item-pq.Top() {
			item -= pq.Sum
		} else {
			item -= (pq.Sum * ((item - pq.Top()) / pq.Sum))
		}
		pq.Push(item)
	}
	return true
}

func main() {}
