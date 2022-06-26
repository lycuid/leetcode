// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
package main

func maxScore(arr []int, k int) int {
	var sum, minSum, windowSum int
	window := len(arr) - k
	for i := 0; i < window; i++ {
		windowSum = arr[i]
	}
	sum, minSum = windowSum, windowSum
	for i := 1; i <= k; i++ {
		windowSum += arr[i+window-1] - arr[i-1]
		if windowSum < minSum {
			minSum = windowSum
		}
		sum += arr[i+window-1]
	}
	return sum - minSum
}

func main() {}
