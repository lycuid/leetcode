// https://leetcode.com/problems/set-mismatch/
package main

func findErrorNums(nums []int) (ret []int) {
	nums, ret = append([]int{0}, nums...), []int{0, 0}
	for i := 1; i < len(nums); i++ {
		for ret[0] != nums[i] && nums[i] != i {
			if nums[i] == nums[nums[i]] {
				ret[0] = nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := 1; ret[1] == 0 && i < len(nums); i++ {
		if nums[i] == ret[0] && i != ret[0] {
			ret[1] = i
		}
	}
	return ret
}

func main() {}
