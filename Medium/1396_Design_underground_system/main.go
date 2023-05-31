// https://leetcode.com/problems/design-underground-system/
package main

import "fmt"

type Tuple struct {
	Station string
	Time    int
}

type UndergroundSystem struct {
	Checkin     map[int]Tuple
	AverageTime map[string]*[2]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		Checkin:     make(map[int]Tuple),
		AverageTime: make(map[string]*[2]int),
	}
}

func Dist(key1, key2 string) string {
	return fmt.Sprintf("%s-%s", key1, key2)
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.Checkin[id] = Tuple{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkin := this.Checkin[id]
	dist, time := Dist(checkin.Station, stationName), t-checkin.Time
	if value, found := this.AverageTime[dist]; found {
		value[0], value[1] = value[0]+time, value[1]+1
	} else {
		this.AverageTime[dist] = &[2]int{time, 1}
	}
	delete(this.Checkin, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	if value, found := this.AverageTime[Dist(startStation, endStation)]; found {
		return float64(value[0]) / float64(value[1])
	}
	return 0
}

func main() {}
