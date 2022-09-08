// https://leetcode.com/problems/queue-reconstruction-by-height/
package main

import "sort"

func reconstructQueue(arr [][]int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] > arr[j][0]
	})
	for i := range arr {
		for j := i; j > 0 && arr[j][1] < j; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}

func main() {}
