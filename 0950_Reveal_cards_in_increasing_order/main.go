// https://leetcode.com/problems/reveal-cards-in-increasing-order/
package main

import "sort"

type Queue struct {
	inner      []int
	start, end int
}

func MakeQ(n int) Queue     { return Queue{make([]int, n), -1, -1} }
func (q Queue) Empty() bool { return q.start == -1 && q.end == -1 }

func (q *Queue) Push(item int) {
	if q.start == -1 && q.end == -1 {
		q.start, q.end = 0, 0
	}
	q.inner[q.end], q.end = item, (q.end+1)%len(q.inner)
}

func (q *Queue) Pop() (item int) {
	item, q.start = q.inner[q.start], (q.start+1)%len(q.inner)
	if q.start == q.end {
		q.start, q.end = -1, -1
	}
	return item
}

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	ret, q := make([]int, len(deck)), Queue{inner: make([]int, len(deck))}
	for i := range ret {
		q.Push(i)
	}
	for !q.Empty() && len(deck) > 0 {
		if ret[q.Pop()], deck = deck[0], deck[1:]; !q.Empty() {
			q.Push(q.Pop())
		}
	}
	return ret
}

func main() {}
