// https://leetcode.com/problems/lexicographical-numbers/
package main

func Aux(num, last int) (nums []int) {
	if num != 0 && num <= last {
		nums = append(nums, num)
		if ns := Aux(num*10, last); len(ns) > 0 {
			nums = append(nums, ns...)
		}
	}
	for i := 1; i <= 9 && num+i <= last; i++ {
		nums = append(nums, num+i)
		if ns := Aux((num+i)*10, last); len(ns) > 0 {
			nums = append(nums, ns...)
		}
	}
	return nums
}

func lexicalOrder(n int) []int {
	return Aux(0, n)
}

func main() {}
