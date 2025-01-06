// https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/
package main

func minOperations(boxes string) []int {
	ret := make([]int, len(boxes))
	left, right := make([]int, len(boxes)+1), make([]int, len(boxes)+1)

	var ballCount int
	for i := 0; i < len(boxes); i++ {
		left[i+1] = ballCount + left[i]
		ballCount += int(boxes[i] - '0')
	}

	ballCount = 0
	for i := len(boxes) - 1; i >= 0; i-- {
		right[i] = ballCount + right[i+1]
		ballCount += int(boxes[i] - '0')
	}
	for i := range ret {
		ret[i] = left[i+1] + right[i]
	}

	return ret
}

func main() {}
