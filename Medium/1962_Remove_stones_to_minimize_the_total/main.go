// https://leetcode.com/problems/remove-stones-to-minimize-the-total/
package main

type PriorityQueue struct{ inner []int }

func MakePQ() PriorityQueue { return PriorityQueue{make([]int, 1)} }

func (this *PriorityQueue) Push(item int) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i] > this.inner[i/2]; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() int {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.inner[j+1] > this.inner[j] {
			j++
		}
		if this.inner[i] > this.inner[j] {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func minStoneSum(piles []int, k int) (ret int) {
	pq := MakePQ()
	for _, num := range piles {
		pq.Push(num)
		ret += num
	}
	for ; k > 0; k-- {
		item := pq.Pop()
		ret -= item / 2
		pq.Push(item - item/2)
	}
	return ret
}

func main() {}
