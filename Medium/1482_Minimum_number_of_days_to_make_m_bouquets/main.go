// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
package main

func minDays(bloomDay []int, m int, k int) int {
	l, r, ret := 0, 0, -1
	for _, day := range bloomDay {
		if day > r {
			r = day
		}
	}
	for l <= r {
		mid := (l + r) / 2
		cons, bs := 0, 0
		for _, day := range bloomDay {
			if day <= mid {
				if cons++; cons < k {
					continue
				}
				bs++
			}
			cons = 0
		}
		if bs >= m {
			r, ret = mid-1, mid
		} else {
			l = mid + 1
		}
	}
	return ret
}

func main() {}
