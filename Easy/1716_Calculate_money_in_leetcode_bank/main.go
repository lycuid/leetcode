// https://leetcode.com/problems/calculate-money-in-leetcode-bank/
package main

func totalMoney(n int) (count int) {
	e := n / 7
	for l, r := 0, 7; l < e; l, r = l+1, r+1 {
		count += (r - l) * (r + l + 1) / 2
	}
	l, r := e, n%7+e
	return count + (r-l)*(r+l+1)/2
}

func main() {}
