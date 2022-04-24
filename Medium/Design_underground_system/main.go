// https://leetcode.com/problems/design-underground-system/
package main

import "fmt"

type Timing struct {
	time, count int
}

type Checkin struct {
	time    int
	station string
}

type UndergroundSystem struct {
	timings  map[string]Timing
	checkins [1e6 + 1]Checkin
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{timings: make(map[string]Timing)}
}

func (this *UndergroundSystem) CheckIn(id int, s string, t int) {
	this.checkins[id] = Checkin{t, s}
}

func (this *UndergroundSystem) CheckOut(id int, checkout string, t int) {
	checkin := this.checkins[id]
	timing := this.timings[Hash(checkin.station, checkout)]
	timing.time += t - checkin.time
	timing.count++
	this.timings[Hash(checkin.station, checkout)] = timing
}

func (this *UndergroundSystem) GetAverageTime(s string, e string) float64 {
	timing := this.timings[Hash(s, e)]
	return float64(timing.time) / float64(timing.count)
}

func Hash(s1, s2 string) string {
	return fmt.Sprintf("%s->%s", s1, s2)
}

func main() {}
