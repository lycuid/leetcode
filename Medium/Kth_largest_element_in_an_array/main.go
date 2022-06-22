// https://leetcode.com/problems/kth-largest-element-in-an-array/
package main

func Heapify(root int, nums []int) {
	if root*2 >= len(nums) {
		return
	}
	child := root * 2
	if root*2+1 < len(nums) && nums[child] < nums[root*2+1] {
		child = root*2 + 1
	}
	if nums[root] < nums[child] {
		nums[child], nums[root] = nums[root], nums[child]
		Heapify(child, nums)
	}
}

type PriorityQueue struct {
	inner []int
}

func MakePQ(n []int) PriorityQueue {
	inner := make([]int, len(n)+1)
	for i, val := range n {
		inner[i+1] = val
	}
	for i := len(inner) / 2; i > 0; i-- {
		Heapify(i, inner)
	}
	return PriorityQueue{inner}
}

func (this *PriorityQueue) Pop() int {
	item := this.inner[1]
	this.inner[1] = this.inner[len(this.inner)-1]
	this.inner = this.inner[:len(this.inner)-1]
	for i := 1; i*2 < len(this.inner); {
		j := i * 2
		if i*2+1 < len(this.inner) && this.inner[i*2+1] > this.inner[i*2] {
			j = i*2 + 1
		}
		if this.inner[i] > this.inner[j] {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
		i = j
	}
	return item
}

func findKthLargest(nums []int, k int) int {
	pq := MakePQ(nums)
	for ; k > 1; k-- {
		pq.Pop()
	}
	return pq.Pop()
}

func main() {}
