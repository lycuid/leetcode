// https://leetcode.com/problems/find-all-duplicates-in-an-array/
package main

func findDuplicates(nums []int) (ret []int) {
	for i := range nums {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			ret = append(ret, num)
		}
	}
	return ret
}

func main() {}
