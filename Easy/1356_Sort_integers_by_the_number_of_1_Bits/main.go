// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
package main

import "sort"

func Zeroes(n int) (count int) {
	for ; n > 0; n >>= 1 {
		count += n & 1
	}
	return count
}

func sortByBits(arr []int) (ret []int) {
	var buckets [32][]int
	for _, n := range arr {
		index := Zeroes(n)
		buckets[index] = append(buckets[index], n)
	}
	for _, arr := range buckets {
		sort.Ints(arr)
		ret = append(ret, arr...)
	}
	return ret
}

func main() {}
