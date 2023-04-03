// https://leetcode.com/problems/boats-to-save-people/
package main

import "sort"

func numRescueBoats(people []int, limit int) (l int) {
	sort.Slice(people, func(i, j int) bool { return people[i] > people[j] })
	for r := len(people) - 1; l <= r; l++ {
		if people[l]+people[r] <= limit {
			r--
		}
	}
	return l
}

func main() {}
