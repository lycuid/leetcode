// https://leetcode.com/problems/minimize-deviation-in-array/
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

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minimumDeviation(nums []int) (ret int) {
	pq, min := MakePQ(), int(1e9+7)
	for _, num := range nums {
		if num&1 == 1 {
			num *= 2
		}
		min = Min(min, num)
		pq.Push(num)
	}
	for ret = 1e9 + 7; ; {
		max := pq.Pop()
		if ret = Min(ret, max-min); max&1 == 1 {
			break
		}
		min = Min(min, max>>1)
		pq.Push(max >> 1)
	}
	return ret
}

func main() {}
