// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/
package main

func findSpecialInteger(arr []int) int {
	for i, n, q := 0, len(arr), len(arr)/4; i < n-q; i++ {
		if arr[i] == arr[i+q] {
			return arr[i]
		}
	}
	return 0
}

func main() {}
