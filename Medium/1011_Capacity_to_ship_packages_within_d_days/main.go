// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
package main

func Days(weights []int, max_capacity int) int {
	var capacity, days int
	for _, weight := range weights {
		if capacity+weight > max_capacity {
			capacity, days = 0, days+1
		}
		capacity += weight
	}
	return days + 1
}

func shipWithinDays(weights []int, days int) int {
	var min, max int
	for _, num := range weights {
		if max += num; num > min {
			min = num
		}
	}
	for min < max {
		if capacity := (min + max) / 2; Days(weights, capacity) <= days {
			max = capacity
		} else {
			min = capacity + 1
		}
	}
	return min
}

func main() {}
