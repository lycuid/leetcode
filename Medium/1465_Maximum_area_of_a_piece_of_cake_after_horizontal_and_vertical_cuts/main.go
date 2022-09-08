// https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
package main

import "sort"

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArea(r int, c int, xs []int, ys []int) int {
	sort.Ints(xs)
	sort.Ints(ys)
	xs, ys = append(xs, r), append(ys, c)
	h, w := xs[0], ys[0]
	for i := 0; i < len(xs)-1; i++ {
		h = Max(h, xs[i+1]-xs[i])
	}
	for i := 0; i < len(ys)-1; i++ {
		w = Max(w, ys[i+1]-ys[i])
	}
	return (h * w) % (1e9 + 7)
}

func main() {}
