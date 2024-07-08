// https://leetcode.com/problems/find-the-winner-of-the-circular-game/
package main

// https://en.wikipedia.org/wiki/Josephus_problem
func findTheWinner(n int, k int) int {
	if n == 1 {
		return 1
	}
	return ((findTheWinner(n-1, k) + k - 1) % n) + 1
}

func main() {}
