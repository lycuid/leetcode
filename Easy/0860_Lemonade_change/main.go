// https://leetcode.com/problems/lemonade-change/
package main

func lemonadeChange(bills []int) bool {
	var change [2]int
	for _, bill := range bills {
		switch bill {
		case 10:
			change[0]++
		case 5:
			change[1]++
		}
		bill -= 5
		for i := range change {
			d := 5 * (len(change) - i)
			if n := bill / d; n <= change[i] {
				change[i] -= n
				bill %= d
			}
		}
		if bill > 0 {
			return false
		}
	}
	return true
}

func main() {}
