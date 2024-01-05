// https://leetcode.com/problems/longest-increasing-subsequence/
package main

func lengthOfLIS(nums []int) int {
	lis := []int{nums[0]}
	for _, num := range nums {
		if l, r := 0, len(lis)-1; lis[r] < num {
			lis = append(lis, num)
		} else {
			for l < r {
				if mid := (l + r) / 2; lis[mid] > num {
					r = mid
				} else if lis[mid] < num {
					l = mid + 1
				} else {
					l = mid
					break
				}
			}
			lis[l] = num
		}
	}
	return len(lis)
}

func main() {}
