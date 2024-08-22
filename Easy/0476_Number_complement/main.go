// https://leetcode.com/problems/number-complement/
package main

func findComplement(num int) (comp int) {
	for i := 0; num > 0; i, num = i+1, num>>1 {
		comp |= ((num & 1) ^ 1) << i
	}
	return comp
}

func main() {}
