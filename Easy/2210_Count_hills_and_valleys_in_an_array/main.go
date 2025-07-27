// https://leetcode.com/problems/count-hills-and-valleys-in-an-array/
package main

func countHillValley(nums []int) (count int) {
	stack := make([]int, 0, len(nums))
	for _, num := range nums {
		for stack = append(stack, num); len(stack) > 1; stack = stack[:len(stack)-1] {
			if n := len(stack); stack[n-1] != stack[n-2] {
				break
			}
		}
	}
	for i := 0; i+2 < len(stack); i++ {
		if n := len(stack); n >= 3 {
			if (stack[i+1] < stack[i] && stack[i+1] < stack[i+2]) ||
				(stack[i+1] > stack[i] && stack[i+1] > stack[i+2]) {
				count++
			}
		}
	}
	return count
}

func main() {}
