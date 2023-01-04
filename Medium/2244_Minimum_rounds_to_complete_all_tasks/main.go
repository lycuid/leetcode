// https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/
package main

func Aux(n int) int {
	if n == 0 {
		return 0
	}
	for _, m := range []int{3, 2} {
		for i := n / m; i > 0; i-- {
			if num := Aux(n - i*m); num != -1 {
				return i + num
			}
		}
	}
	return -1
}

func minimumRounds(tasks []int) (ret int) {
	count := make(map[int]int)
	for _, task := range tasks {
		count[task]++
	}
	for _, v := range count {
		if v = Aux(v); v == -1 {
			return -1
		}
		ret += v
	}
	return ret
}

func main() {}
