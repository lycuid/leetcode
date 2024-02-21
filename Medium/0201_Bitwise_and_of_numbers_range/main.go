// https://leetcode.com/problems/bitwise-and-of-numbers-range/
package main

func rangeBitwiseAnd(left int, right int) (ret int) {
	for i := 0; left > 0 || right > 0; i++ {
		if l, r := left&1, right&1; l^r == 0 {
			ret += l & r << i
		} else {
			ret = 0
		}
		left, right = left>>1, right>>1
	}
	return ret
}

func main() {}
