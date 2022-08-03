// https://leetcode.com/problems/my-calendar-i/
package main

type Interval struct {
	start, end int
}

func (this *Interval) LessThan(interval Interval) bool {
	return this.end <= interval.start
}

type MyCalendar struct {
	inner []Interval
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	interval, i := Interval{start, end}, 0
	for ; i < len(this.inner); i++ {
		if interval.LessThan(this.inner[i]) {
			break
		} else if !this.inner[i].LessThan(interval) {
			return false
		}
	}
	this.inner = append(this.inner, interval)
	for j := len(this.inner) - 1; j > i; j-- {
		this.inner[j], this.inner[j-1] = this.inner[j-1], this.inner[j]
	}
	return true
}

func main() {}
