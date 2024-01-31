// https://leetcode.com/problems/daily-temperatures/
package main

func dailyTemperatures(temperatures []int) []int {
	stack, ret := []int{}, make([]int, len(temperatures))
	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			ret[index], stack = i-index, stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}

func main() {}
