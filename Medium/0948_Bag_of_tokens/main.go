// https://leetcode.com/problems/bag-of-tokens/
package main

import "sort"

func bagOfTokensScore(tokens []int, power int) (ret int) {
	sort.Ints(tokens)
	for i, j := 0, len(tokens)-1; i <= j; {
		if tokens[i] > power && i < j && ret > 0 {
			power, j, ret = power+tokens[j], j-1, ret-1
		} else if tokens[i] <= power {
			power, i, ret = power-tokens[i], i+1, ret+1
		} else {
			break
		}
	}
	return ret
}

func main() {}
