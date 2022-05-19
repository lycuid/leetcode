// https://leetcode.com/problems/shuffle-the-array/
package main

func shuffle(nums []int, n int) (ret []int) {
	for i := 0; i < n; i++ {
		ret = append(ret, nums[i])
		ret = append(ret, nums[n+i])
	}
	return ret
}

func main() {}
