// https://leetcode.com/problems/daily-temperatures/
package main

func dailyTemperatures(temperatures []int) []int {
	ret, stack := make([]int, len(temperatures)), []int{}
	for i, temp := range temperatures {
		for n := len(stack) - 1; n >= 0 && temp > temperatures[stack[n]]; n = len(stack) - 1 {
			ret[stack[n]], stack = i-stack[n], stack[:n]
		}
		stack = append(stack, i)
	}
	return ret
}

func main() {}
