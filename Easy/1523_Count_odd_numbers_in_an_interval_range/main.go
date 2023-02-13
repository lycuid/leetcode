// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
package main

func countOdds(l int, h int) int { return (h-l+1)>>1 + (l%2)*(h%2) }

func main() {}
