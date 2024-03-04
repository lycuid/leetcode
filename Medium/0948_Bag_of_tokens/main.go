// https://leetcode.com/problems/bag-of-tokens/
package main

import "sort"

func bagOfTokensScore(tokens []int, power int) (score int) {
	sort.Ints(tokens)
	for i, j := 0, len(tokens)-1; i <= j; {
		if power < tokens[i] && i < j && score > 0 {
			power, j, score = power+tokens[j], j-1, score-1
		} else if power >= tokens[i] {
			power, i, score = power-tokens[i], i+1, score+1
		} else {
			break
		}
	}
	return score
}

func main() {}
