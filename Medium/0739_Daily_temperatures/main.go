// https://leetcode.com/problems/daily-temperatures/
package main

func dailyTemperatures(temp []int) []int {
	ret, stack, cursor := make([]int, len(temp)), make([]int, len(temp)), 0
	for i := range temp {
		for ; cursor > 0 && temp[stack[cursor-1]] < temp[i]; cursor-- {
			ret[stack[cursor-1]] = i - stack[cursor-1]
		}
		stack[cursor], cursor = i, cursor+1
	}
	return ret
}

func main() {}
