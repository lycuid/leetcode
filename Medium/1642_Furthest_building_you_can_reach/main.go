// https://leetcode.com/problems/furthest-building-you-can-reach/
package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type PQ struct {
	inner  []int
	cursor int
}

func MakePQ(n int) PQ { return PQ{make([]int, n+2), 1} }

func (this *PQ) Push(item int) {
	this.inner[this.cursor] = item
	for i := this.cursor; i > 1 && this.inner[i/2] > this.inner[i]; i /= 2 {
		this.inner[i/2], this.inner[i] = this.inner[i], this.inner[i/2]
	}
	this.cursor++
}

func (this *PQ) Pop() (item int) {
	item, this.cursor = this.inner[1], this.cursor-1
	this.inner[1] = this.inner[this.cursor]
	for i := 1; i*2 < this.cursor; {
		j := i * 2
		if i*2+1 < this.cursor && this.inner[i*2+1] < this.inner[j] {
			j = i*2 + 1
		}
		if this.inner[i] < this.inner[j] {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
		i = j
	}
	this.inner[this.cursor] = -1
	return item
}

func furthestBuilding(heights []int, bricks int, ladders int) (index int) {
	if ladders >= len(heights) {
		return len(heights) - 1
	}
	pq := MakePQ(ladders + 1)
	for i := 1; i < ladders+1; i++ {
		pq.Push(Max(0, heights[i]-heights[i-1]))
	}
	for index = ladders; index < len(heights)-1 && bricks >= 0; index++ {
		pq.Push(Max(0, heights[index+1]-heights[index]))
		bricks -= pq.Pop()
	}
	if bricks < 0 {
		index--
	}
	return index
}

func main() {}
