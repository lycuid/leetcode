// https://leetcode.com/problems/number-of-matching-subsequences/
package main

func NextGreater(num int, nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < num {
			l = mid + 1
		} else if nums[mid] > num {
			r = mid
		} else {
			if mid < len(nums)-1 {
				return nums[mid+1]
			}
			l, r = mid, mid
		}
	}
	if l < len(nums) && nums[l] > num {
		return nums[l]
	}
	return num
}

func numMatchingSubseq(s string, words []string) (ret int) {
	var indices [26][]int
	for i, ch := range s {
		indices[ch-'a'] = append(indices[ch-'a'], i)
	}
OUTTER:
	for _, word := range words {
		current := -1
		for _, ch := range word {
			tmp := NextGreater(current, indices[ch-'a'])
			if tmp <= current {
				continue OUTTER
			}
			current = tmp
		}
		ret++
	}
	return ret
}

func main() {}
