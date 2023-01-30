// https://leetcode.com/problems/n-th-tribonacci-number/
package main

func tribonacci(n int) int {
	tri := [3]int{0, 1, 1}
	if n <= 2 {
		return tri[n]
	}
	for i := 3; i <= n; i++ {
		tri[0], tri[1], tri[2] = tri[1], tri[2], tri[0]+tri[1]+tri[2]
	}
	return tri[2]
}

func main() {}
