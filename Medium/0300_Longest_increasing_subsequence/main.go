// https://leetcode.com/problems/longest-increasing-subsequence/
package main

func lengthOfLIS(nums []int) (ret int) {
	lis := []int{nums[0]}
	for _, num := range nums {
		if l, r := 0, len(lis)-1; lis[r] >= num {
			for mid := (l + r) / 2; l < r; mid = (l + r) / 2 {
				if lis[mid] > num {
					r = mid
				} else if lis[mid] < num {
					l = mid + 1
				} else {
					l = mid
					break
				}
			}
			lis[l] = num
		} else {
			lis = append(lis, num)
		}
	}
	return len(lis)
}

func main() {}
