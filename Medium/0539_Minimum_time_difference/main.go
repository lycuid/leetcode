// https://leetcode.com/problems/minimum-time-difference/
package main

import (
	"sort"
	"time"
)

func findMinDifference(timePoints []string) int {
	cache, min_duration := make([]time.Time, len(timePoints)), 24*time.Hour
	for i, str := range timePoints {
		cache[i], _ = time.Parse("15:04", str)
	}
	sort.Slice(cache, func(i, j int) bool {
		return cache[i].Before(cache[j])
	})
	cache = append(cache, cache[0].Add(24*time.Hour))
	for i := 1; i < len(cache); i++ {
		if duration := cache[i].Sub(cache[i-1]); duration < min_duration {
			min_duration = duration
		}
	}
	return int(min_duration.Minutes())
}

func main() {
}
