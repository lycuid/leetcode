// https://leetcode.com/problems/calculate-money-in-leetcode-bank/
package main

func totalMoney(n int) (count int) {
	for n--; n >= 0; n-- {
		count += n/7 + n%7 + 1
	}
	return count
}

func main() {}
