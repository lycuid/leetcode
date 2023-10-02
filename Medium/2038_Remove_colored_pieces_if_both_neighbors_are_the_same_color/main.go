// https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
package main

func winnerOfGame(colors string) bool {
	var score [2]int
	for i, j := 0, 0; i < len(colors); {
		current := colors[i]
		for j = i; i < len(colors) && colors[i] == current; i++ {
		}
		if n := i - j - 2; n > 0 {
			score[current-'A'] += n
		}
	}
	return score[0] > score[1]
}

func main() {}
