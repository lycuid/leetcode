// https://leetcode.com/problems/number-of-1-bits/
package main

func hammingWeight(num uint32) (ones int) {
	for i := 0; i < 32; i++ {
		ones += int((num >> i) & 1)
	}
	return ones
}

func main() {}
