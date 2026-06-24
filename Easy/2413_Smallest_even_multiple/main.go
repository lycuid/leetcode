// https://leetcode.com/problems/smallest-even-multiple/
package main

func smallestEvenMultiple(n int) int {
	return n + n*(n&1)
}

func main() {}
