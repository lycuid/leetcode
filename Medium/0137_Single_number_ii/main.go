// https://leetcode.com/problems/single-number-ii/
package main

func singleNumber(nums []int) int {
	var ret int32
	for i := 0; i < 32; i++ {
		count := 0
		for _, num := range nums {
			count += (num >> i) & 1
		}
		if count%3 != 0 {
			ret += 1 << i
		}
	}
	return int(ret)
}

func main() {}
