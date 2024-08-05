// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/
package main

type PriorityQueue struct{ items []*Stream }

func MakePQ() PriorityQueue { return PriorityQueue{make([]*Stream, 1)} }

func (pq PriorityQueue) Less(i, j int) bool { return pq.items[i].GetValue() < pq.items[j].GetValue() }

func (pq *PriorityQueue) Push(item *Stream) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.Less(i, i/2); i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *PriorityQueue) Pop() *Stream {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.Less(j+1, j) {
			j++
		}
		if pq.Less(i, j) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

type Stream struct {
	items      []int
	index, sum int
}

func (this Stream) Done() bool     { return this.index >= len(this.items) }
func (this *Stream) GetValue() int { return this.sum + this.items[this.index] }
func (this *Stream) Advance()      { this.sum, this.index = this.sum+this.items[this.index], this.index+1 }

func rangeSum(nums []int, n int, left int, right int) (ret int) {
	const MOD = 1e9 + 7
	pq := MakePQ()
	for i := range nums {
		pq.Push(&Stream{items: nums, index: i})
	}
	for i := 1; i < left; i++ {
		stream := pq.Pop()
		if stream.Advance(); !stream.Done() {
			pq.Push(stream)
		}
	}
	for i := left; i <= right; i++ {
		stream := pq.Pop()
		ret = (ret + stream.GetValue()) % MOD
		if stream.Advance(); !stream.Done() {
			pq.Push(stream)
		}
	}
	return ret % MOD
}

func main() {}
