// https://leetcode.com/problems/find-k-closest-elements/
package main

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k
	for l < r {
		if mid := (l + r) / 2; x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return arr[l : l+k]
}
