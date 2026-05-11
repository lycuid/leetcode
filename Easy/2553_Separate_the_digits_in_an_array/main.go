// https://leetcode.com/problems/separate-the-digits-in-an-array/
package main

func separateDigits(nums []int) (res []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for ; nums[i] > 0; nums[i] /= 10 {
			res = append(res, nums[i]%10)
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func main() {}
