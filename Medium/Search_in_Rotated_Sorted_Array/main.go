// https://leetcode.com/problems/search-in-rotated-sorted-array/

package main

func get_pivot(nums []int) (pivot int) {
  low, high := 0, len(nums) - 1
  for low < high {
    if low == high - 1 {
      if nums[low] < nums[high] {
        return low
      } else {
        return high
      }
    }

    pivot = (low + high) / 2

    if nums[pivot] < nums[low] {
      high = pivot
    } else if nums[pivot] > nums[high] {
      low = pivot
      // if none of the above is true, that means the array is not
      // rotated.
    } else { return 0 }
  }
  return pivot
}

func search(nums []int, target int) int {
  pivot := get_pivot(nums)

  low, high := 0, len(nums) - 1
  if pivot > 0 {
    if target <= nums[high] {
      low = pivot
    } else if target >= nums[low] {
      high = pivot - 1
    }
  }

  for low <= high {
    mid := (low + high) / 2
    if target > nums[mid] {
      low = mid + 1
    } else if target < nums[mid] {
      high = mid - 1
    } else {
      return mid
    }
  }
  return -1
}

func main() { }

