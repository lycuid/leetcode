// https://leetcode.com/problems/koko-eating-bananas/
package main

func Hours(k int, piles []int) (hours int) {
	for i := range piles {
		if hours += piles[i] / k; piles[i]%k != 0 {
			hours++
		}
	}
	return hours
}

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for i := range piles {
		if r < piles[i] {
			r = piles[i]
		}
	}
	for l < r {
		if k := (l + r) / 2; Hours(k, piles) <= h {
			r = k
		} else {
			l = k + 1
		}
	}
	return l
}

func main() {}
