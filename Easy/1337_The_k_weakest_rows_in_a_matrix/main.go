// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/description/?envType=daily-question&envId=2023-09-18
package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	for i := range mat {
		for j := 1; j < len(mat[i]); j++ {
			mat[i][0] += mat[i][j]
		}
		mat[i][1] = i
	}
	sort.Slice(mat, func(i, j int) bool {
		if mat[i][0] == mat[j][0] {
			return mat[i][1] < mat[j][1]
		}
		return mat[i][0] < mat[j][0]
	})
	cache := make([]int, k)
	for i := range cache {
		cache[i] = mat[i][1]
	}
	return cache[:k]
}

func main() {}
