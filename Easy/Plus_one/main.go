// https://leetcode.com/problems/plus-one/

package main

func plusOne(digits []int) []int {
  carry := 1
  i := len(digits) - 1
  for i >= 0 && carry > 0 {
    digits[i] += carry
    carry = int(digits[i] / 10)
    digits[i] %= 10
    i--
  }
  if carry > 0 {
    digits = append([]int{carry}, digits...)
  }
  return digits
}

func main() { }

