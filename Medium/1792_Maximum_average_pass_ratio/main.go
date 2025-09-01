// https://leetcode.com/problems/maximum-average-pass-ratio/
package main

type PriorityQueue struct{ items [][]float64 }

func MakePQ() PriorityQueue {
	return PriorityQueue{make([][]float64, 1)}
}

func (pq PriorityQueue) Empty() bool {
	return len(pq.items) <= 1
}

func (pq PriorityQueue) less(i, j int) bool {
	n := (pq.items[i][0]+1)/(pq.items[i][1]+1) - pq.items[i][0]/pq.items[i][1]
	m := (pq.items[j][0]+1)/(pq.items[j][1]+1) - pq.items[j][0]/pq.items[j][1]
	return n > m
}

func (pq *PriorityQueue) Push(item []float64) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.less(i, i/2); i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *PriorityQueue) Pop() []float64 {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.less(j+1, j) {
			j++
		}
		if pq.less(i, j) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

func maxAverageRatio(classes [][]int, extraStudents int) (count float64) {
	pq := MakePQ()
	for _, class := range classes {
		pq.Push([]float64{float64(class[0]), float64(class[1])})
	}
	for ; extraStudents > 0; extraStudents-- {
		class := pq.Pop()
		class[0], class[1] = class[0]+1, class[1]+1
		pq.Push(class)
	}
	for !pq.Empty() {
		class := pq.Pop()
		count += class[0] / class[1]
	}
	return count / float64(len(classes))
}

func main() {}
