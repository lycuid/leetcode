// https://leetcode.com/problems/number-of-1-bits/
package main

func hammingWeight(num uint32) (count int) {
	for ; num > 0; num >>= 1 {
		count += int(num & 1)
	}
	return count
}

func main() {}
