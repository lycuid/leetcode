// https://leetcode.com/problems/maximum-subsequence-score/
package main

import "sort"

type Tuple struct{ num1, num2 int }

type PriorityQueue struct{ inner []Tuple }

func MakePQ() PriorityQueue {
	return PriorityQueue{make([]Tuple, 1)}
}

func (this *PriorityQueue) Push(item Tuple) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i].num1 < this.inner[i/2].num1; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() Tuple {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.inner[j+1].num1 < this.inner[j].num1 {
			j++
		}
		if this.inner[i].num1 < this.inner[j].num1 {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func maxScore(nums1 []int, nums2 []int, k int) (ret int64) {
	cache, pq, sum := make([]Tuple, len(nums1)), MakePQ(), 0
	for i := 0; i < len(nums1); i++ {
		cache[i] = Tuple{nums1[i], nums2[i]}
	}
	sort.Slice(cache, func(i, j int) bool {
		return cache[i].num2 > cache[j].num2
	})
	for i := 0; i < k-1; i++ {
		pq.Push(cache[0])
		sum, cache = sum+cache[0].num1, cache[1:]
	}
	pq.Push(cache[0])
	sum = sum + cache[0].num1
	ret = int64(sum * cache[0].num2)
	for cache = cache[1:]; len(cache) > 0; cache = cache[1:] {
		prev, curr := pq.Pop(), cache[0]
		sum += curr.num1 - prev.num1
		if n := int64(sum * curr.num2); n > ret {
			ret = n
		}
		pq.Push(curr)
	}
	return ret
}

func main() {}
