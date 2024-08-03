// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/
package main

func canBeEqual(target []int, arr []int) bool {
	Reverse := func(start, end int) {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	for i := range arr {
		j := i
		for j < len(arr) && arr[j] != target[i] {
			j++
		}
		if j >= len(arr) {
			return false
		}
		Reverse(i, j)
	}
	return true
}

func main() {}
