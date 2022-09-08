// https://leetcode.com/problems/powx-n/

package main

func myPow(x float64, n int) float64 {
  if n == 0 {
    return 1.0
  }
  if n < 0 {
    n = 0 - n
    x = 1 / x
  }
  num := myPow(x, n/2)
  if n % 2 == 0 {
    return num * num
  } else {
    return num * num * x
  }
}

func main() { }

