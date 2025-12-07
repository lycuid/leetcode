// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
package main

func countOdds(low int, high int) int {
	return (low & 1) + ((high - low - high&1) >> 1) + (high & 1)
}

func main() {}
