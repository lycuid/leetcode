// https://leetcode.com/problems/1-bit-and-2-bit-characters/
package main

func isOneBitCharacter(bits []int) bool {
	for ; len(bits) > 1; bits = bits[1:] {
		bits = bits[bits[0]:]
	}
	return len(bits) == 1
}

func main() {}
