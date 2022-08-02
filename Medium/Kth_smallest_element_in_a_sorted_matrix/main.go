// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
package main

type PriorityQueue struct {
	inner [][]int
}

func MakePQ() PriorityQueue {
	return PriorityQueue{make([][]int, 1)}
}

func (this *PriorityQueue) Push(item []int) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i][0] < this.inner[i/2][0]; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() []int {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.inner[j+1][0] < this.inner[j][0] {
			j++
		}
		if this.inner[i][0] < this.inner[j][0] {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func kthSmallest(matrix [][]int, k int) int {
	pq := MakePQ()
	for _, arr := range matrix {
		pq.Push(arr)
	}
	for i := 0; i < k-1; i++ {
		if arr := pq.Pop(); len(arr) > 1 {
			pq.Push(arr[1:])
		}
	}
	return pq.Pop()[0]
}

func main() {}
