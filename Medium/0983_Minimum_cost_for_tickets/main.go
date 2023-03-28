// https://leetcode.com/problems/minimum-cost-for-tickets/
package main

type Solution struct {
	days, costs, cache []int
	size               int
}

func MakeSolution(days, costs []int) Solution {
	return Solution{
		days:  days,
		costs: costs,
		cache: make([]int, len(days)+1),
		size:  len(days),
	}
}

func (this *Solution) Solve(start int) int {
	if start != this.size && this.cache[start] == 0 {
		pos := [3]int{start + 1, start, start}
		for pos[1] < this.size && this.days[pos[1]] < this.days[start]+7 {
			pos[1]++
		}
		for pos[2] < this.size && this.days[pos[2]] < this.days[start]+30 {
			pos[2]++
		}
		this.cache[start] = Min(
			this.costs[0]+this.Solve(pos[0]),
			this.costs[1]+this.Solve(pos[1]),
			this.costs[2]+this.Solve(pos[2]),
		)
	}
	return this.cache[start]
}

func Min(num int, nums ...int) int {
	for i := range nums {
		if nums[i] < num {
			num = nums[i]
		}
	}
	return num
}

func mincostTickets(days []int, costs []int) int {
	solution := MakeSolution(days, costs)
	return solution.Solve(0)
}

func main() {}
