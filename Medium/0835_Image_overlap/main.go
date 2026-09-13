// https://leetcode.com/problems/image-overlap/
package main

import "math/bits"

func largestOverlap(img1 [][]int, img2 [][]int) (res int) {
	n := len(img1)
	img := [2][]uint{make([]uint, n), make([]uint, n)}
	for i, nums := range img1 {
		for _, num := range nums {
			img[0][i] = img[0][i]<<1 | uint(num)
		}
	}
	for i, nums := range img2 {
		for _, num := range nums {
			img[1][i] = img[1][i]<<1 | uint(num)
		}
	}
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			var tl, bl, tr, br int
			for j := i; j < n; j++ {
				tl += bits.OnesCount(img[0][j] << k & img[1][j-i])
				bl += bits.OnesCount(img[0][j-i] << k & img[1][j])
				tr += bits.OnesCount(img[0][j] >> k & img[1][j-i])
				br += bits.OnesCount(img[0][j-i] >> k & img[1][j])
			}
			res = max(res, tl, bl, tr, br)
		}
	}
	return res
}

func main() {}
