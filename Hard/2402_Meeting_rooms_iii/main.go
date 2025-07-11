// https://leetcode.com/problems/meeting-rooms-iii/
package main

import "sort"

type priorityQueue struct {
	items []any
	less  func(any, any) bool
}

func makePQ(less func(any, any) bool) priorityQueue {
	return priorityQueue{make([]any, 1), less}
}

func (pq *priorityQueue) push(item any) {
	pq.items = append(pq.items, item)
	for i := len(pq.items) - 1; i > 1 && pq.less(pq.items[i], pq.items[i/2]); i /= 2 {
		pq.items[i], pq.items[i/2] = pq.items[i/2], pq.items[i]
	}
}

func (pq *priorityQueue) pop() any {
	item, n := pq.items[1], len(pq.items)-1
	pq.items[1], pq.items = pq.items[n], pq.items[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.less(pq.items[j+1], pq.items[j]) {
			j++
		}
		if pq.less(pq.items[i], pq.items[j]) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	}
	return item
}

func (pq priorityQueue) top() *any {
	if len(pq.items) == 1 {
		return nil
	}
	return &pq.items[1]
}

func mostBooked(n int, meetings [][]int) (room int) {
	if len(meetings) > 0 {
		score := make([]int, n)

		type Slot struct{ room, end int }
		slots := makePQ(func(i, j any) bool {
			if i.(Slot).end == j.(Slot).end {
				return i.(Slot).room < j.(Slot).room
			}
			return i.(Slot).end < j.(Slot).end
		})

		rooms := makePQ(func(i, j any) bool {
			return i.(int) < j.(int)
		})
		for i := 0; i < n; i++ {
			rooms.push(i)
		}

		sort.Slice(meetings, func(i, j int) bool {
			return meetings[i][0] < meetings[j][0]
		})

		for _, meeting := range meetings {
			for slots.top() != nil && (*slots.top()).(Slot).end <= meeting[0] {
				rooms.push(slots.pop().(Slot).room)
			}

			var slot Slot
			if rooms.top() != nil {
				slot.room = rooms.pop().(int)
			} else {
				slot = slots.pop().(Slot)
				meeting[1] += slot.end - meeting[0]
			}
			slot.end = meeting[1]
			score[slot.room]++
			slots.push(slot)
		}

		for i := range score {
			if score[i] > score[room] {
				room = i
			}
		}
	}
	return room
}

func main() {}
