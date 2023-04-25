// https://leetcode.com/problems/smallest-number-in-infinite-set/
package main

type PriorityQueue struct{ inner []int }

func MakePQ() PriorityQueue            { return PriorityQueue{make([]int, 1)} }
func (this PriorityQueue) Empty() bool { return len(this.inner) == 1 }

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

type SmallestInfiniteSet struct {
	pq      PriorityQueue
	has     map[int]bool
	current int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{MakePQ(), make(map[int]bool), 1}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.pq.Empty() {
		this.current++
		return this.current - 1
	}
	num := this.pq.Pop()
	delete(this.has, num)
	return num
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num == this.current-1 {
		this.current--
	} else if num < this.current && !this.has[num] {
		this.has[num] = true
		this.pq.Push(num)
	}
}

func main() {}
