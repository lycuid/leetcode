// https://leetcode.com/problems/power-of-four/
package main

func isPowerOfFour(n int) (ret bool) {
	if ret = n > 0 && n&(n-1) == 0; ret {
		for ; n > 1; n >>= 1 {
			ret = !ret
		}
	}
	return ret
}

func main() {}
