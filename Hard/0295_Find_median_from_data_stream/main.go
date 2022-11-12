// https://leetcode.com/problems/find-median-from-data-stream/
package main

type Comparator = func(int, int) bool

type PriorityQueue struct {
	inner []int
	Less  Comparator
}

func MakePQ(l Comparator) PriorityQueue { return PriorityQueue{make([]int, 1), l} }
func (this PriorityQueue) Size() int    { return len(this.inner) - 1 }
func (this PriorityQueue) Top() int     { return this.inner[1] }

func (this *PriorityQueue) Push(item int) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.Less(this.inner[i], this.inner[i/2]); i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() int {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.Less(this.inner[j+1], this.inner[j]) {
			j++
		}
		if this.Less(this.inner[i], this.inner[j]) {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

type MedianFinder struct{ upper, lower PriorityQueue }

func Constructor() MedianFinder {
	return MedianFinder{
		upper: MakePQ(func(i, j int) bool { return i < j }),
		lower: MakePQ(func(i, j int) bool { return i > j }),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.lower.Push(num)
	for this.upper.Push(this.lower.Pop()); this.upper.Size() > this.lower.Size(); {
		this.lower.Push(this.upper.Pop())
	}
}

func (this *MedianFinder) FindMedian() (mid float64) {
	if mid = float64(this.lower.Top()); this.upper.Size() == this.lower.Size() {
		mid = (mid + float64(this.upper.Top())) / 2
	}
	return mid
}

func main() {}
