// https://leetcode.com/problems/minimize-maximum-of-array/
package main

func minimizeArrayValue(nums []int) (ret int) {
	for i, sum := 0, 0; i < len(nums); i++ {
		sum += nums[i]
		if n := (sum + i) / (i + 1); n > ret {
			ret = n
		}
	}
	return ret
}

func main() {}
