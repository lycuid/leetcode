// https://leetcode.com/problems/snapshot-array/
package main

import "math"

type Tuple struct{ snap_id, value int }

type SnapshotArray struct {
	snap_id int
	inner   [][]Tuple
}

func Constructor(length int) SnapshotArray {
	inner := make([][]Tuple, length)
	for i := range inner {
		inner[i] = []Tuple{{}, {math.MaxInt, 0}}
	}
	return SnapshotArray{0, inner}
}

func (this *SnapshotArray) Set(index int, val int) {
	n := len(this.inner[index]) - 2
	if this.inner[index][n].snap_id != this.snap_id {
		this.inner[index][n+1] = Tuple{snap_id: this.snap_id}
		this.inner[index] = append(this.inner[index], Tuple{snap_id: math.MaxInt})
		n++
	}
	this.inner[index][n].value = val
}

func (this *SnapshotArray) Snap() int {
	this.snap_id++
	return this.snap_id - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	l, r := 0, len(this.inner[index])-1
	for l < r {
		if mid := (l + r) / 2; this.inner[index][mid].snap_id > snap_id {
			r = mid
		} else if this.inner[index][mid].snap_id < snap_id {
			l = mid + 1
		} else {
			l, r = mid, mid
		}
	}
	for this.inner[index][l].snap_id > snap_id {
		l--
	}
	return this.inner[index][l].value
}

func main() {}
