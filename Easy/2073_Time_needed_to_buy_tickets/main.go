// https://leetcode.com/problems/time-needed-to-buy-tickets/
package main

func timeRequiredToBuy(tickets []int, k int) (ret int) {
	for i, val := range tickets {
		if i > k && val > tickets[k]-1 {
			val = tickets[k] - 1
		}
		if val > tickets[k] {
			val = tickets[k]
		}
		ret += val
	}
	return ret
}

func main() {}
