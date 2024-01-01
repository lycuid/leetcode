// https://leetcode.com/problems/assign-cookies/
package main

type PriorityQueue struct{ inner []int }

func MakePQ(inner []int) (pq PriorityQueue) {
	pq.inner = append([]int{0}, inner...)
	for i := len(pq.inner) / 2; i > 0; i-- {
		Heapify(pq.inner, i)
	}
	return pq
}

func Heapify(xs []int, x int) {
	max := x
	if l := 2 * x; l < len(xs) && xs[l] > xs[max] {
		max = l
	}
	if r := 2*x + 1; r < len(xs) && xs[r] > xs[max] {
		max = r
	}
	if max != x {
		xs[x], xs[max] = xs[max], xs[x]
		Heapify(xs, max)
	}
}

func (pq PriorityQueue) Empty() bool { return len(pq.inner) == 1 }
func (pq PriorityQueue) Peek() int   { return pq.inner[1] }

func (pq *PriorityQueue) Pop() int {
	item, n := pq.inner[1], len(pq.inner)-1
	pq.inner[1], pq.inner = pq.inner[n], pq.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.inner[j+1] > pq.inner[j] {
			j++
		}
		if pq.inner[i] > pq.inner[j] {
			break
		}
		pq.inner[i], pq.inner[j] = pq.inner[j], pq.inner[i]
	}
	return item
}

func findContentChildren(g []int, s []int) (count int) {
	pq_g, pq_s := MakePQ(g), MakePQ(s)
	for ; !pq_g.Empty() && !pq_s.Empty(); pq_g.Pop() {
		if pq_g.Peek() <= pq_s.Peek() {
			count, _ = count+1, pq_s.Pop()
		}
	}
	return count
}

func main() {}
