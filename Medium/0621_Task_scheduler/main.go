// https://leetcode.com/problems/task-scheduler
package main

type Tuple struct{ ch, count, cooldown int }

type PriorityQueue struct{ inner []Tuple }

func MakePQ() PriorityQueue          { return PriorityQueue{make([]Tuple, 1)} }
func (pq PriorityQueue) Empty() bool { return len(pq.inner) == 1 }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq.inner[i].count == pq.inner[j].count {
		return pq.inner[i].ch < pq.inner[j].ch
	}
	return pq.inner[i].count > pq.inner[j].count
}

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

func leastInterval(tasks []byte, n int) (cycles int) {
	var freq [26]int
	for _, ch := range tasks {
		freq[ch-'A']++
	}
	pq, stash := MakePQ(), []Tuple{}
	for i := range freq {
		if freq[i] > 0 {
			pq.Push(Tuple{i, freq[i], 0})
		}
	}
	for ; !pq.Empty() || len(stash) > 0; cycles++ {
		if !pq.Empty() {
			if tup := pq.Pop(); tup.count > 1 {
				tup.cooldown, tup.count = n+1, tup.count-1
				stash = append(stash, tup)
			}
		}
		for i := range stash {
			stash[i].cooldown--
		}
		if len(stash) > 0 && stash[0].cooldown == 0 {
			pq.Push(stash[0])
			stash = stash[1:]
		}
	}
	return cycles
}

func main() {}
