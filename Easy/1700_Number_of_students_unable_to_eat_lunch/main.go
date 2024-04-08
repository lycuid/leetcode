// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
package main

func countStudents(students []int, sandwiches []int) int {
	cache, ret := [2]int{}, len(students)
	for _, num := range students {
		cache[num]++
	}
	for i := 0; i < len(sandwiches) && cache[sandwiches[i]] > 0; i++ {
		cache[sandwiches[i]], ret = cache[sandwiches[i]]-1, ret-1
	}
	return ret
}

func main() {}
