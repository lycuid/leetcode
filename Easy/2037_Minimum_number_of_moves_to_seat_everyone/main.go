// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
package main

import "sort"

func minMovesToSeat(seats []int, students []int) (count int) {
	Abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	sort.Ints(seats)
	sort.Ints(students)
	for i := range seats {
		count += Abs(seats[i] - students[i])
	}
	return count
}

func main() {}
