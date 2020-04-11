// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3284/

package main

func toDigits(n int) (ns []int) {
  for n > 0 {
    ns = append([]int{n % 10}, ns...)
    n /= 10
  }
  return ns
}

func isOne(nums []int) bool {
  s := 0
  for _, val := range nums {
    s += val
    if s > 1 {
      return false
    }
  }
  return true 
}

func isHappy(n int) bool {
  seen := make(map[int]bool)
  num := toDigits(n)

  for !isOne(num) {
    m := 0
    for _, m1 := range num {
      m += (m1 * m1)
    }
    if _, ok := seen[m]; ok {
      return false
    }
    seen[m] = true
    num = toDigits(m)
  }
  
  return true
}

func main() { }
