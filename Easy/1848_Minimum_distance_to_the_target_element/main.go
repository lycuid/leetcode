// https://leetcode.com/problems/minimum-distance-to-the-target-element/
package main

func getMinDistance(nums []int, target int, start int) (i int) {
	for ; start+i < len(nums) || start-i >= 0; i++ {
		if (start+i < len(nums) && nums[start+i] == target) || (start-i >= 0 && nums[start-i] == target) {
			break
		}
	}
	return i
}

func main() {}
