// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/
package main

func minBitFlips(start int, goal int) (count int) {
	for ; start+goal > 0; start, goal = start>>1, goal>>1 {
		count += start&1 ^ goal&1
	}
	return count
}

func main() {}
