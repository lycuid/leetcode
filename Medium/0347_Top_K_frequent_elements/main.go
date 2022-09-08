// https://leetcode.com/problems/top-k-frequent-elements/
package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func topKFrequent(nums []int, k int) []int {
	max_freq := nums[0]
	freq := make(map[int]int) // a[n] = count
	for _, val := range nums {
		freq[val] += 1
		max_freq = Max(max_freq, freq[val])
	}
	weight_at := make([][]int, max_freq+1) // a[count] = {n1, n2, ..}
	for n, count := range freq {
		weight_at[count] = append(weight_at[count], n)
	}
	ret := make([]int, k)
	for i := max_freq; i >= 0 && k > 0; i-- {
		for _, val := range weight_at[i] {
			ret[k-1] = val
			k--
		}
	}
	return ret
}

func main() {}
