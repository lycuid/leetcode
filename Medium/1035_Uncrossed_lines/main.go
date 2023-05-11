// https://leetcode.com/problems/uncrossed-lines/
package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n := len(nums2)
	cache := [2][]int{make([]int, n+1), make([]int, n+1)}
	for _, num := range nums1 {
		for i, j := 1, 1; j <= n; j++ {
			if nums2[j-1] == num {
				cache[i][j] = cache[i-1][j-1] + 1
			} else if cache[i-1][j] > cache[i][j-1] {
				cache[i][j] = cache[i-1][j]
			} else {
				cache[i][j] = cache[i][j-1]
			}
		}
		cache[0], cache[1] = cache[1], cache[0]
	}
	return cache[0][n]
}

func main() {}
