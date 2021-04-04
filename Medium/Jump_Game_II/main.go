// https://leetcode.com/problems/jump-game-ii/

package main

import "math"

func minint(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func jump(nums []int) int {
	l := len(nums)
	weight := make([]int, l)

	for i := 0; i < l; i++ {
		weight[i] = int(math.Pow10(6))
	}
	weight[0] = 0

	for i, num := range nums {
		for j := i; j < minint(i+num+1, l); j++ {
			weight[j] = minint(weight[j], weight[i]+1)
		}
	}

	return weight[l-1]
}

func main() {}
