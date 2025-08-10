// https://leetcode.com/problems/reordered-power-of-2/
package main

import (
	"slices"
	"sort"
)

func reorderedPowerOf2(n int) bool {
	digits := func(n int) (d []int) {
		for ; n > 0; n /= 10 {
			d = append(d, n%10)
		}
		sort.Ints(d)
		return d
	}
	numDigits := digits(n)
	for i := 0; ; i++ {
		finalDigits := digits(1 << i)
		if len(finalDigits) > len(numDigits) {
			break
		}
		if eq := slices.Equal(finalDigits, numDigits); eq {
			return true
		}
	}
	return false
}

func main() {}
