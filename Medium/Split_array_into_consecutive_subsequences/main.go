// https://leetcode.com/problems/split-array-into-consecutive-subsequences/
package main

type Group struct {
	number, count int
}

func isPossible(nums []int) bool {
	var groups []Group
	for _, num := range nums {
		index, count := -1, len(nums)+1
		for i, group := range groups {
			if group.number == num-1 && group.count < count {
				index, count = i, group.count
			}
		}
		if index == -1 {
			groups = append(groups, Group{num, 1})
		} else {
			groups[index].count++
			groups[index].number = num
		}
	}
	for _, group := range groups {
		if group.count < 3 {
			return false
		}
	}
	return true
}

func main() {}
