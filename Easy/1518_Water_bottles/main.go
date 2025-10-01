// https://leetcode.com/problems/water-bottles/
package main

func numWaterBottles(numBottles int, numExchange int) (count int) {
	for empty := 0; numBottles > 0; {
		count += numBottles
		numBottles, empty = (numBottles+empty)/numExchange, (numBottles+empty)%numExchange
	}
	return count
}

func main() {}
