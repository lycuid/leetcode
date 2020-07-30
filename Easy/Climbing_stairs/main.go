// https://leetcode.com/problems/climbing-stairs/

package main

func climbStairs(n int) int {
  p := 0
  p2 := 1
  for n >= 2 {
    p2 += p
    p = p2 - p
    n--
  }
  return p + p2
}

func main() { }

