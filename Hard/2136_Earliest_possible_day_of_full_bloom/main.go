// https://leetcode.com/problems/earliest-possible-day-of-full-bloom/
package main

import "sort"

type Tuple struct {
	x, y int
}

func earliestFullBloom(plantTime []int, growTime []int) (ret int) {
	tuples, time := make([]Tuple, len(plantTime)), 0
	for i := range plantTime {
		tuples = append(tuples, Tuple{plantTime[i], growTime[i]})
	}
	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].y > tuples[j].y
	})
	for _, tuple := range tuples {
		time += tuple.x
		if n := time + tuple.y; n > ret {
			ret = n
		}
	}
	return ret
}

func main() {}
