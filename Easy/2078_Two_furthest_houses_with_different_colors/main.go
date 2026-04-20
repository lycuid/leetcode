// https://leetcode.com/problems/two-furthest-houses-with-different-colors/
package main

func maxDistance(colors []int) (res int) {
	for i := range colors {
		for j := i - 1; j >= 0; j-- {
			if colors[i] != colors[j] {
				res = max(res, i-j)
			}
		}
	}
	return res
}

func main() {}
