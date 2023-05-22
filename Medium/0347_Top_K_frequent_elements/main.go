// https://leetcode.com/problems/top-k-frequent-elements/
package main

func topKFrequent(nums []int, k int) []int {
	freq, ret, max := make(map[int]int), make([]int, k), 0
	for _, num := range nums {
		if freq[num]++; freq[num] > max {
			max = freq[num]
		}
	}
	bucket := make([][]int, max+1)
	for k, v := range freq {
		bucket[v] = append(bucket[v], k)
	}
	for i := len(bucket) - 1; k > 0 && i >= 0; i-- {
		for j := 0; k > 0 && j < len(bucket[i]); j++ {
			ret[k-1], k = bucket[i][j], k-1
		}
	}
	return ret
}

func main() {}
