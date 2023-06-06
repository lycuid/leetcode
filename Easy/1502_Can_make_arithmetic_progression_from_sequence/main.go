// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i, dx := 1, arr[1]-arr[0]; i < len(arr); i++ {
		if arr[i]-arr[i-1] != dx {
			return false
		}
	}
	return true
}

func main() {}
