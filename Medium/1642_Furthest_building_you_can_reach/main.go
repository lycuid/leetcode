// https://leetcode.com/problems/furthest-building-you-can-reach/
package main

type PriorityQueue struct{ inner []int }

func MakePQ() PriorityQueue          { return PriorityQueue{make([]int, 1)} }
func (pq PriorityQueue) Empty() bool { return len(pq.inner) == 1 }

func (pq *PriorityQueue) Push(item int) {
	pq.inner = append(pq.inner, item)
	for i := len(pq.inner) - 1; i > 1 && pq.inner[i] < pq.inner[i/2]; i /= 2 {
		pq.inner[i], pq.inner[i/2] = pq.inner[i/2], pq.inner[i]
	}
}

func (pq *PriorityQueue) Pop() int {
	item, n := pq.inner[1], len(pq.inner)-1
	pq.inner[1], pq.inner = pq.inner[n], pq.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.inner[j+1] < pq.inner[j] {
			j++
		}
		if pq.inner[i] < pq.inner[j] {
			break
		}
		pq.inner[i], pq.inner[j] = pq.inner[j], pq.inner[i]
	}
	return item
}

func furthestBuilding(heights []int, bricks int, ladders int) (index int) {
	if n := len(heights); ladders >= n {
		return n - 1
	}
	Max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	pq := MakePQ()
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
