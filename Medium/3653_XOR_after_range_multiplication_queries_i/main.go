// https://leetcode.com/problems/xor-after-range-multiplication-queries-i/
package main

func xorAfterQueries(nums []int, queries [][]int) (num int) {
	for _, q := range queries {
		for i := q[0]; i <= q[1]; i += q[2] {
			nums[i] = (nums[i] * q[3]) % (1e9 + 7)
		}
	}
	for i := range nums {
		num ^= nums[i]
	}
	return num
}

func main() {}
