// https://leetcode.com/problems/minimum-time-to-complete-trips/
package main

import "math"

func Trips(n int, time []int) (ntrips int) {
	for i := range time {
		if time[i] <= n {
			ntrips += n / time[i]
		}
	}
	return ntrips
}

func minimumTime(time []int, totalTrips int) int64 {
	l, r := 0, math.MaxInt
	for i := range time {
		if time[i]*totalTrips < r {
			r = time[i] * totalTrips
		}
	}
	for l < r {
		if mid := (l + r) / 2; Trips(mid, time) >= totalTrips {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}

func main() {}
