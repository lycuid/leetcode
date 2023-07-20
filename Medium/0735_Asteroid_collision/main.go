// https://leetcode.com/problems/asteroid-collision/
package main

func asteroidCollision(asteroids []int) []int {
	stack, cursor := make([]int, len(asteroids)), 0
	for _, num := range asteroids {
		for cursor > 0 && stack[cursor-1] > 0 && stack[cursor-1] < -num {
			cursor--
		}
		if cursor != 0 && stack[cursor-1] > 0 && num < 0 {
			if stack[cursor-1] == -num {
				cursor--
			}
		} else {
			stack[cursor], cursor = num, cursor+1
		}
	}
	return stack[:cursor]
}

func main() {}
