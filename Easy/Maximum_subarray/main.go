// https://leetcode.com/problems/maximum-subarray/

package main

func maxSubArray(nums []int) int {
  if len(nums) == 0 {
    return 0
  }
  total := nums[0]
  for i := 1; i < len(nums); i++ {
    newInt := nums[i] + nums[i - 1]
    if newInt > nums[i] {
      nums[i] = newInt
    }
    if nums[i] > total {
      total = nums[i]
    }
  }
  return total
}

func main() { }
