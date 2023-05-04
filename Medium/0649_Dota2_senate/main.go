// https://leetcode.com/problems/dota2-senate/
package main

type Ring struct {
	inner      []int
	start, end int
}

func MakeRing(n int) Ring     { return Ring{make([]int, n+1), 0, 0} }
func (this Ring) Empty() bool { return this.start == this.end }

func (this *Ring) Push(item int) {
	this.inner[this.end%len(this.inner)] = item
	this.end = (this.end + 1) % len(this.inner)
}

func (this *Ring) Pop() int {
	item := this.inner[this.start%len(this.inner)]
	this.start = (this.start + 1) % len(this.inner)
	return item
}

func predictPartyVictory(senate string) string {
	r, d := MakeRing(len(senate)), MakeRing(len(senate))
	for i := range senate {
		if senate[i] == 'R' {
			r.Push(i)
		} else {
			d.Push(i)
		}
	}
	for n := len(senate); !r.Empty() && !d.Empty(); {
		ri, di := r.Pop(), d.Pop()
		if ri < di {
			r.Push(ri + n)
		} else {
			d.Push(di + n)
		}
	}
	if r.Empty() {
		return "Dire"
	}
	return "Radiant"
}

func main() {}
