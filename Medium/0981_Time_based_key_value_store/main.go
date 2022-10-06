// https://leetcode.com/problems/time-based-key-value-store/
package main

type Tuple struct {
	time  int
	value string
}

type TimeMap struct {
	inner map[string][]Tuple
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]Tuple)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.inner[key] = append(this.inner[key], Tuple{timestamp, value})
}

func (this TimeMap) Get(key string, timestamp int) (ret string) {
	if l, r := 0, len(this.inner[key])-1; r >= 0 && this.inner[key][l].time <= timestamp {
		for l < r {
			if mid := (l + r) / 2; this.inner[key][mid].time < timestamp {
				l = mid + 1
			} else if this.inner[key][mid].time > timestamp {
				r = mid
			} else {
				l, r = mid, mid
			}
		}
		for ; l > 0 && this.inner[key][l].time > timestamp; l-- {
		}
		ret = this.inner[key][l].value
	}
	return ret
}

func main() {}
