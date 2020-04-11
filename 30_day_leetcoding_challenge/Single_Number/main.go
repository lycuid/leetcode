// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3283/

package main

func singleNumber(nums []int) int {
  num := nums[0]
  for _, n := range nums[1:] {
    num ^= n
  }
  return num
}

func main() { }

