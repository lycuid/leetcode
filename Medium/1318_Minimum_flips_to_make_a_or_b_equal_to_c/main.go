// https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
package main

func Count(ab, c int) int {
	if c == 1 {
		if ab == 0 {
			return c
		} else {
			return 0
		}
	}
	return ab
}

func minFlips(a int, b int, c int) (ret int) {
	for ; c > 0; a, b, c = a>>1, b>>1, c>>1 {
		ret += Count((a&1)+(b&1), c&1)
	}
	for ; a > 0 || b > 0; a, b = a>>1, b>>1 {
		ret += Count((a&1)+(b&1), 0)
	}
	return ret
}

func main() {}
