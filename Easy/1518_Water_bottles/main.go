// https://leetcode.com/problems/water-bottles/
package main

func numWaterBottles(numBottles int, numExchange int) int {
	count := numBottles
	for numBottles >= numExchange {
		exchanged := numBottles / numExchange
		count += exchanged
		numBottles = numBottles%numExchange + exchanged
	}
	return count
}

func main() {}
