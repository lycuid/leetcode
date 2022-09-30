// https://leetcode.com/problems/the-skyline-problem/
package main

import "sort"

type PriorityQueue struct {
	inner []int
}

func (this PriorityQueue) Top() (item int) { return this.inner[1] }
func (this PriorityQueue) Empty() bool     { return len(this.inner) == 1 }

func (this *PriorityQueue) Add(item int) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i] > this.inner[i/2]; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Remove(item int) {
	n := len(this.inner) - 1
	for index := 1; index < n; index++ {
		if this.inner[index] == item {
			this.inner[index] = this.inner[n]
			for i, j := index, index*2; j < n; i, j = j, j*2 {
				if j+1 < n && this.inner[j+1] > this.inner[j] {
					j++
				}
				if this.inner[i] > this.inner[j] {
					break
				}
				this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
			}
			break
		}
	}
	if n >= 1 {
		this.inner = this.inner[:n]
	}
}

func getSkyline(buildings [][]int) (ret [][]int) {
	var pairs [][2]int
	for _, b := range buildings {
		pairs = append(pairs, [2]int{b[0], -b[2]})
		pairs = append(pairs, [2]int{b[1], b[2]})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
	hs, h := PriorityQueue{make([]int, 2)}, 0
	for _, p := range pairs {
		if p[1] < 0 {
			hs.Add(-p[1])
		} else {
			hs.Remove(p[1])
		}
		if !hs.Empty() && h != hs.Top() {
			h, ret = hs.Top(), append(ret, []int{p[0], hs.Top()})
		}
	}
	return ret
}

func main() {}
