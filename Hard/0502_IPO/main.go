// https://leetcode.com/problems/ipo/
package main

type Tuple struct{ capital, profit int }
type Comparator = func(t1, t2 Tuple) bool

type PriorityQueue struct {
	inner []Tuple
	cmp   Comparator
}

func MakePQ(cmp Comparator) PriorityQueue {
	return PriorityQueue{make([]Tuple, 1), cmp}
}

func (pq PriorityQueue) Top() *Tuple {
	if len(pq.inner) > 1 {
		return &pq.inner[1]
	}
	return nil
}

func (pq *PriorityQueue) Push(item Tuple) {
	pq.inner = append(pq.inner, item)
	for i := len(pq.inner) - 1; i > 1 && pq.cmp(pq.inner[i], pq.inner[i/2]); i /= 2 {
		pq.inner[i], pq.inner[i/2] = pq.inner[i/2], pq.inner[i]
	}
}

func (pq *PriorityQueue) Pop() Tuple {
	item, n := pq.inner[1], len(pq.inner)-1
	pq.inner[1], pq.inner = pq.inner[n], pq.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.cmp(pq.inner[j+1], pq.inner[j]) {
			j++
		}
		if pq.cmp(pq.inner[i], pq.inner[j]) {
			break
		}
		pq.inner[i], pq.inner[j] = pq.inner[j], pq.inner[i]
	}
	return item
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	min_pq := MakePQ(func(t1, t2 Tuple) bool { return t1.capital < t2.capital })
	max_pq := MakePQ(func(t1, t2 Tuple) bool { return t1.profit > t2.profit })
	for i := range profits {
		min_pq.Push(Tuple{capital: capital[i], profit: profits[i]})
	}
	for ; k > 0; k-- {
		for min_pq.Top() != nil && min_pq.Top().capital <= w {
			max_pq.Push(min_pq.Pop())
		}
		if max_pq.Top() == nil {
			break
		}
		w += max_pq.Pop().profit
	}
	return w
}

func main() {}
