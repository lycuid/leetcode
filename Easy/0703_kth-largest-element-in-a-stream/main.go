// https://leetcode.com/problems/kth-largest-element-in-a-stream/
package main

import "sort"

const Inf = -(1 << 31)

type KthLargest struct {
	k      int // constraints: 1 <= k <= 10^4
	stream []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	stream := make([]int, k)
	for i := 0; i < k; i++ {
		stream[i] = Inf
	}
	ret := KthLargest{k: k, stream: stream}
	for _, val := range nums {
		ret.Add(val)
	}
	return ret
}

func (this *KthLargest) Add(val int) int {
	if val > this.stream[0] {
		this.stream[0] = val
		for i := 0; i < this.k-1; i++ {
			if this.stream[i] <= this.stream[i+1] {
				break
			}
			this.stream[i], this.stream[i+1] = this.stream[i+1], this.stream[i]
		}
	}
	index := 0
	for ; index < this.k && this.stream[index] == Inf; index++ {
	}
	if index == this.k {
		panic("I find the lack of index, Disturbing!..")
	}
	return this.stream[index]
}

func main() { }
