// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
package main

func sumZero(n int) []int {
	if n%2 == 1 {
		return append(sumZero(n-1), 0)
	}
	ret := make([]int, 0, n)
	for i := 1; i <= n/2; i++ {
		ret = append(ret, i, -i)
	}
	return ret
}

func main() {}
