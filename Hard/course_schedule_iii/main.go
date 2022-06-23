// https://leetcode.com/problems/course-schedule-iii/
package main

import "sort"

type PriorityQueue struct {
	inner []int
}

func MakePQ() PriorityQueue {
	return PriorityQueue{make([]int, 1)}
}

func (this *PriorityQueue) Insert(item int) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.inner[i/2] < this.inner[i]; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Remove() int {
	item, n := this.inner[1], len(this.inner)
	this.inner[1], this.inner = this.inner[n-1], this.inner[:n-1]
	for i, j := 1, 2; j < n-1; i, j = j, j*2 {
		if j+1 < len(this.inner) && this.inner[j+1] > this.inner[j] {
			j++
		}
		if this.inner[i] > this.inner[j] {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func (this PriorityQueue) Size() int {
	return len(this.inner) - 1
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	pq, start := MakePQ(), 0
	for _, c := range courses {
		pq.Insert(c[0])
		start += c[0]
		for start > c[1] {
			start -= pq.Remove()
		}
	}
	return pq.Size()
}

func main() {}
