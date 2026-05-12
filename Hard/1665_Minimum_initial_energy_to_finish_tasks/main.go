// https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/
package main

import "sort"

func minimumEffort(tasks [][]int) (res int) {
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i][1]-tasks[i][0] == tasks[j][1]-tasks[j][0] {
			return tasks[i][1] > tasks[j][1]
		}
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})
	var carry int
	for _, task := range tasks {
		res += max(0, task[1]-carry)
		carry = max(carry, task[1]) - task[0]
	}
	return res
}

func main() {}
