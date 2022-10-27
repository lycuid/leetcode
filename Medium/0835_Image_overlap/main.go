// https://leetcode.com/problems/image-overlap/
package main

import "math/bits"

func F(n, x1, x2, bit int) uint {
	return uint(bit&((1<<(n-x1))-1)) >> (n - x2 - 1)
}

func Aux(n, x1, x2, y1, y2 int, bit [2][]int) (ret int) {
	dy, dx := y2-y1, x2-x1
	for y := 0; y+dy < n; y++ {
		for x := 0; x+dx < n; x++ {
			t := 0
			for i := 0; i <= dy; i++ {
				t += bits.OnesCount(
					F(n, x1, x2, bit[0][y1+i]) & F(n, x, x+dx, bit[1][y+i]),
				)
			}
			if t > ret {
				ret = t
			}
		}
	}
	return ret
}

func largestOverlap(img1 [][]int, img2 [][]int) (ret int) {
	var bit [2][]int
	n := len(img1)
	prefix := make([][]int, n+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := range img1 {
		for j := range img1[i] {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + img1[i][j]
		}
	}
	for i := range bit {
		bit[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			bit[0][i] = bit[0][i]<<1 | img1[i][j]
			bit[1][i] = bit[1][i]<<1 | img2[i][j]
		}
	}
	for y1 := 0; y1 < n; y1++ {
		for y2 := n - 1; y2 >= y1; y2-- {
			for x1 := 0; x1 < n; x1++ {
				for x2 := n - 1; x2 >= x1; x2-- {
					if prefix[y2+1][x2+1]-prefix[y1][x1] > ret {
						if max := Aux(n, x1, x2, y1, y2, bit); max > ret {
							ret = max
						}
					}
				}
			}
		}
	}
	return ret
}

func main() {}
