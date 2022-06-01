// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
package main

func smallerNumbersThanCurrent(nums []int) (ret []int) {
	var numbers [102]int
	ret = make([]int, len(nums))
	for _, n := range nums {
		numbers[n+1]++
	}
	for i := 1; i <= 101; i++ {
		numbers[i] += numbers[i-1]
	}
	for i := range nums {
		ret[i] = numbers[nums[i]]
	}
	return ret
}

func main() {}
