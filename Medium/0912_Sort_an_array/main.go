// https://leetcode.com/problems/sort-an-array/
package main

func merge(nums []int, start, mid, end int) {
	i, j, tmp := start, mid+1, []int{}
	for i <= mid && j <= end {
		if nums[i] < nums[j] {
			tmp, i = append(tmp, nums[i]), i+1
		} else {
			tmp, j = append(tmp, nums[j]), j+1
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := range tmp {
		nums[start+i] = tmp[i]
	}
}

func mergesort(nums []int, start, end int) {
	if mid := (start + end) / 2; start < end {
		mergesort(nums, start, mid)
		mergesort(nums, mid+1, end)
		merge(nums, start, mid, end)
	}
}

func sortArray(nums []int) []int {
	mergesort(nums, 0, len(nums)-1)
	return nums
}

func main() {}
