// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
package main

func findWinners(matches [][]int) [][]int {
	var losses [1e5 + 1]int
	for i := range losses {
		losses[i] = -1
	}
	for _, match := range matches {
		if losses[match[0]] == -1 {
			losses[match[0]] = 0
		}
		if losses[match[1]] == -1 {
			losses[match[1]] = 0
		}
		losses[match[1]]++
	}
	var zeroes, ones []int
	for player, loss := range losses {
		if loss == 0 {
			zeroes = append(zeroes, player)
		}
		if loss == 1 {
			ones = append(ones, player)
		}
	}
	return [][]int{zeroes, ones}
}

func main() {}
