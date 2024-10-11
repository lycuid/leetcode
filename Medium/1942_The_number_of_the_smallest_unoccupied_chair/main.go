// https://leetcode.com/problems/the-number-of-the-smallest-unoccupied-chair/
package main

import "sort"

type PriorityQueue struct{ items []int }

func MakePQ() PriorityQueue { return PriorityQueue{make([]int, 1)} }

func (pq PriorityQueue) Empty() bool { return len(pq.items) == 1 }

func (pq *PriorityQueue) Push(item int) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.items[i] < pq.items[i/2]; i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *PriorityQueue) Pop() int {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.items[j+1] < pq.items[j] {
			j++
		}
		if pq.items[i] < pq.items[j] {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

func smallestChair(times [][]int, targetFriend int) int {
	var (
		arrival   = make([]int, len(times))
		departure = make([]int, len(times))
		chair     = make([]int, len(times))
	)
	for i := range arrival {
		arrival[i] = i
	}
	for i := range departure {
		departure[i] = i
	}
	sort.Slice(arrival, func(i, j int) bool {
		return times[arrival[i]][0] < times[arrival[j]][0]
	})
	sort.Slice(departure, func(i, j int) bool {
		return times[departure[i]][1] < times[departure[j]][1]
	})

	new_chair, pq := 1, MakePQ()
	for _, in := range arrival {
		for ; len(departure) > 0 && times[departure[0]][1] <= times[in][0]; departure = departure[1:] {
			pq.Push(chair[departure[0]])
		}
		if !pq.Empty() {
			chair[in] = pq.Pop()
		} else {
			chair[in], new_chair = new_chair, new_chair+1
		}
		if in == targetFriend {
			break
		}
	}
	return chair[targetFriend] - 1
}

func main() {}
