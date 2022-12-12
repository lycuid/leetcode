// https://leetcode.com/problems/climbing-stairs/
package main

func climbStairs(n int) (snd int) {
	for fst := 1; n >= 0; n-- {
		fst, snd = snd, fst+snd
	}
	return snd
}

func main() {}
