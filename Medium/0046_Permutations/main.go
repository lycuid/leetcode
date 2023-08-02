// https://leetcode.com/problems/permutations/
package main

func permute(nums []int) (ret [][]int) {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	head, tail := nums[0], permute(nums[1:])
	for _, lst := range tail {
		for i := 0; i <= len(lst); i++ {
			tmp := make([]int, len(lst[:i]))
			copy(tmp, lst[:i])
			ret = append(ret, append(tmp, append([]int{head}, lst[i:]...)...))
		}
	}
	return ret
}

func main() {}
