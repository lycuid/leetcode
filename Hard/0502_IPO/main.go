// https://leetcode.com/problems/ipo/
package main

type Tuple struct{ fst, snd int }
type Cmp = func(Tuple, Tuple) bool

type PriorityQueue struct {
	inner []Tuple
	F     Cmp
}

func MakePQ(F Cmp) PriorityQueue       { return PriorityQueue{F: F, inner: make([]Tuple, 1)} }
func (this PriorityQueue) Empty() bool { return len(this.inner) == 1 }
func (this PriorityQueue) Top() Tuple  { return this.inner[1] }

func (this *PriorityQueue) Push(item Tuple) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.F(this.inner[i], this.inner[i/2]); i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() Tuple {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.F(this.inner[j+1], this.inner[j]) {
			j++
		}
		if this.F(this.inner[i], this.inner[j]) {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	aux := MakePQ(func(t1, t2 Tuple) bool { return t1.snd < t2.snd })
	for i := range profits {
		aux.Push(Tuple{profits[i], capital[i]})
	}
	for pq := MakePQ(func(t1, t2 Tuple) bool { return t1.fst > t2.fst }); k > 0; k-- {
		for !aux.Empty() && aux.Top().snd <= w {
			pq.Push(aux.Pop())
		}
		if pq.Empty() {
			break
		}
		w += pq.Pop().fst
	}
	return w
}

func main() {}
