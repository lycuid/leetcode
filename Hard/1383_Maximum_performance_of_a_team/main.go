// https://leetcode.com/problems/maximum-performance-of-a-team/
package main

import "sort"

type Tuple struct {
	Speed, Efficiency int
}

type PriorityQueue struct {
	inner []int
}

func MakePQ() PriorityQueue            { return PriorityQueue{make([]int, 1)} }
func (this PriorityQueue) Top() int    { return this.inner[1] }
func (this PriorityQueue) Length() int { return len(this.inner) - 1 }

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

func maxPerformance(n int, speed []int, efficiency []int, k int) (ret int) {
	tuples := make([]Tuple, n)
	for i := range tuples {
		tuples[i] = Tuple{speed[i], efficiency[i]}
	}
	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].Efficiency > tuples[j].Efficiency
	})
	pq, totalSpeed := MakePQ(), 0
	for _, t := range tuples {
		for pq.Length() >= k {
			totalSpeed -= pq.Pop()
		}
		totalSpeed += t.Speed
		pq.Push(t.Speed)
		if performance := totalSpeed * t.Efficiency; ret < performance {
			ret = performance
		}
	}
	return ret % (1e9 + 7)
}

func main() {}
