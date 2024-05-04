// https://leetcode.com/problems/boats-to-save-people/
package main

import "sort"

func numRescueBoats(people []int, limit int) (count int) {
	sort.Ints(people)
	for l, r := 0, len(people)-1; l <= r; r, count = r-1, count+1 {
		if people[l]+people[r] <= limit {
			l++
		}
	}
	return count
}

func main() {}
