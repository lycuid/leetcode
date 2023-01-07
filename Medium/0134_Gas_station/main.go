// https://leetcode.com/problems/gas-station/
package main

func canCompleteCircuit(gas []int, cost []int) (index int) {
	gas, cost = append([]int{0}, gas...), append([]int{0}, cost...)
	for i := 1; i < len(cost); i++ {
		if cost[i] = cost[i-1] + (cost[i] - gas[i]); cost[i] > cost[index] {
			index = i
		}
	}
	if cost[len(cost)-1] > cost[0] {
		index = -1
	}
	return index
}

func main() {}
