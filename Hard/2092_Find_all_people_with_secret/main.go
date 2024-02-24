// https://leetcode.com/problems/find-all-people-with-secret/
package main

type Tuple struct{ person, time int }
type PriorityQueue struct{ inner []Tuple }

func MakePQ() PriorityQueue { return PriorityQueue{make([]Tuple, 1)} }

func (pq PriorityQueue) Empty() bool        { return len(pq.inner) == 1 }
func (pq PriorityQueue) Less(i, j int) bool { return pq.inner[i].time < pq.inner[j].time }

func (pq *PriorityQueue) Push(item Tuple) {
	pq.inner = append(pq.inner, item)
	for i := len(pq.inner) - 1; i > 1 && pq.Less(i, i/2); i /= 2 {
		pq.inner[i], pq.inner[i/2] = pq.inner[i/2], pq.inner[i]
	}
}

func (pq *PriorityQueue) Pop() Tuple {
	item, n := pq.inner[1], len(pq.inner)-1
	pq.inner[1], pq.inner = pq.inner[n], pq.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.Less(j+1, j) {
			j++
		}
		if pq.Less(i, j) {
			break
		}
		pq.inner[i], pq.inner[j] = pq.inner[j], pq.inner[i]
	}
	return item
}

func findAllPeople(n int, meetings [][]int, firstPerson int) (ret []int) {
	var (
		pq    = MakePQ()
		adj   = make(map[int][]Tuple)
		knows = make([]bool, n)
		done  = make([]bool, n)
	)

	knows[0] = true
	meetings = append(meetings, []int{firstPerson, 0, 0})
	for _, meeting := range meetings {
		adj[meeting[0]] = append(adj[meeting[0]], Tuple{meeting[1], meeting[2]})
		adj[meeting[1]] = append(adj[meeting[1]], Tuple{meeting[0], meeting[2]})
	}

	for pq.Push(Tuple{0, 0}); !pq.Empty(); {
		tuple := pq.Pop()
		done[tuple.person] = true
		for _, meeting := range adj[tuple.person] {
			if !done[meeting.person] && meeting.time >= tuple.time {
				pq.Push(meeting)
				knows[meeting.person] = true
			}
		}
	}

	for i := range knows {
		if knows[i] {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {}
