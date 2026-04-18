// https://leetcode.com/problems/mirror-distance-of-an-integer/
package main

func mirrorDistance(n int) int {
	var rev int
	for m := n; m > 0; m /= 10 {
		rev = rev*10 + m%10
	}
	res := n - rev
	if res < 0 {
		res = -res
	}
	return res
}

func main() {}
