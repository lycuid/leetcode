// https://leetcode.com/problems/my-calendar-i/
package main

import "slices"

type MyCalendar struct{ bookings [][2]int }

func Constructor() MyCalendar { return MyCalendar{} }

func (this *MyCalendar) Book(start int, end int) bool {
	i := len(this.bookings) - 1
	for ; i >= 0; i-- {
		if this.bookings[i][0] >= end {
			continue
		}
		if this.bookings[i][1] <= start {
			break
		}
		return false
	}
	this.bookings = slices.Insert(this.bookings, i+1, [2]int{start, end})
	return true
}

func main() {}
