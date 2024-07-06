// https://leetcode.com/problems/pass-the-pillow/
package main

func passThePillow(n int, time int) int {
	d := time % (n - 1)
	if (time/(n-1))&1 == 1 {
		return n - d
	}
	return 1 + d
}

func main() {}
