// https://leetcode.com/problems/my-calendar-iii/
package main

import "sort"

type MyCalendarThree struct {
	events map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{make(map[int]int)}
}

func (this *MyCalendarThree) Book(s int, e int) (k int) {
	this.events[s], this.events[e] = this.events[s]+1, this.events[e]-1
	current, keys := 0, make([]int, len(this.events))
	for key := range this.events {
		keys[current] = key
		current++
	}
	sort.Ints(keys)
	current = 0
	for _, key := range keys {
		if current += this.events[key]; current > k {
			k = current
		}
	}
	return k
}

func main() {}
