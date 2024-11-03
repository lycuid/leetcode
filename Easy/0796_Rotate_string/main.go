// https://leetcode.com/problems/rotate-string/
package main

func rotateString(s string, goal string) bool {
	if n := len(s); n == len(goal) {
	mainloop:
		for i := 0; i < n; i++ {
			for j := range goal {
				if s[(i+j)%n] != goal[j] {
					continue mainloop
				}
			}
			return true
		}
	}
	return false
}

func main() {}
