// https://leetcode.com/problems/peak-index-in-a-mountain-array/
package main

func peakIndexInMountainArray(arr []int) int {
	l, r := 1, len(arr)-2
	for l < r {
		if mid := (l + r) / 2; arr[mid-1] > arr[mid] {
			r = mid - 1
		} else if arr[mid+1] > arr[mid] {
			l = mid + 1
		} else {
			l, r = mid, mid
		}
	}
	return l
}

func main() {}
