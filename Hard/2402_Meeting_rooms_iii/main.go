// https://leetcode.com/problems/meeting-rooms-iii/
package main

import "sort"

type Comparable[T any] interface{ Less(T) bool }
type Room uint64
type Tuple struct{ room, start, end uint64 }

func (this Room) Less(that Room) bool   { return this < that }
func (this Tuple) Less(that Tuple) bool {
	if this.end == that.end {
		return this.room < that.room
	}
	return this.end < that.end
}

type PriorityQueue[T Comparable[T]] struct{ inner []T }

func MakePQ[T Comparable[T]]() PriorityQueue[T] {
	return PriorityQueue[T]{make([]T, 1)}
}

func (pq PriorityQueue[T]) Top() (item *T) {
	if len(pq.inner) > 1 {
		item = &pq.inner[1]
	}
	return item
}

func (pq *PriorityQueue[T]) Push(item T) {
	pq.inner = append(pq.inner, item)
	for i := len(pq.inner) - 1; i > 1 && pq.inner[i].Less(pq.inner[i/2]); i /= 2 {
		pq.inner[i], pq.inner[i/2] = pq.inner[i/2], pq.inner[i]
	}
}

func (pq *PriorityQueue[T]) Pop() T {
	item, n := pq.inner[1], len(pq.inner)-1
	pq.inner[1], pq.inner = pq.inner[n], pq.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.inner[j+1].Less(pq.inner[j]) {
			j++
		}
		if pq.inner[i].Less(pq.inner[j]) {
			break
		}
		pq.inner[i], pq.inner[j] = pq.inner[j], pq.inner[i]
	}
	return item
}

func mostBooked(n int, meetings [][]int) (max int) {
	if n > 1 && len(meetings) > 1 {
		cost, meets, rooms := make([]int, n), MakePQ[Tuple](), MakePQ[Room]()
		for i := 0; i < n; i++ {
			rooms.Push(Room(i))
		}
		sort.Slice(meetings, func(i, j int) bool {
			return meetings[i][0] < meetings[j][0]
		})
		for _, meeting := range meetings {
			for meets.Top() != nil && meets.Top().end <= uint64(meeting[0]) {
				rooms.Push(Room(meets.Pop().room))
			}
			var curr Tuple
			if rooms.Top() != nil {
				curr = Tuple{
					room:  uint64(int(rooms.Pop())),
					start: uint64(meeting[0]),
					end:   uint64(meeting[1]),
				}
			} else {
				prev := meets.Pop()
				curr = Tuple{
					room:  prev.room,
					start: prev.end,
					end:   uint64(meeting[1]) + prev.end - uint64(meeting[0]),
				}
			}
			cost[curr.room]++
			meets.Push(curr)
		}
		for i := len(cost) - 1; i >= 0; i-- {
			if cost[i] >= cost[max] {
				max = i
			}
		}
	}
	return max
}

func main() {}
