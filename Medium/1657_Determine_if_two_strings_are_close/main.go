// https://leetcode.com/problems/determine-if-two-strings-are-close/
package main

func Sorted(nums [26]int) [26]int {
	Sort(0, 25, &nums)
	return nums
}

func Sort(start, end int, nums *[26]int) {
	if mid := (start + end) / 2; start < end {
		Sort(start, mid, nums)
		Sort(mid+1, end, nums)
		Merge(start, mid, end, nums)
	}
}

func Merge(start, mid, end int, nums *[26]int) {
	tmp, i, j, k := make([]int, end-start+1), start, mid+1, 0
	for i <= mid && j <= end {
		if nums[i] < nums[j] {
			tmp[k], i, k = nums[i], i+1, k+1
		} else {
			tmp[k], j, k = nums[j], j+1, k+1
		}
	}
	for i <= mid {
		tmp[k], i, k = nums[i], i+1, k+1
	}
	for j <= end {
		tmp[k], j, k = nums[j], j+1, k+1
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
}

func closeStrings(word1 string, word2 string) bool {
	var freq [2][26]int
	var chars [2][26]bool
	for _, ch := range word1 {
		freq[0][ch-'a'], chars[0][ch-'a'] = freq[0][ch-'a']+1, true
	}
	for _, ch := range word2 {
		freq[1][ch-'a'], chars[1][ch-'a'] = freq[1][ch-'a']+1, true
	}
	return chars[0] == chars[1] && Sorted(freq[0]) == Sorted(freq[1])
}

func main() {}
