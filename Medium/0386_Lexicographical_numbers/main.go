// https://leetcode.com/problems/lexicographical-numbers/
package main

func lexicalOrder(n int) []int {
	var (
		ret = make([]int, 0, n+1)
		Aux func(int)
	)
	Aux = func(num int) {
		if num <= n {
			ret = append(ret, num)
		}
		num *= 10
		for i := 0; i < 10 && num+i <= n; i++ {
			Aux(num + i)
		}
	}
	for i := 1; i < 10 && i <= n; i++ {
		Aux(i)
	}
	return ret
}

func main() {}
