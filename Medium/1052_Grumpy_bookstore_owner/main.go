// https://leetcode.com/problems/grumpy-bookstore-owner/
package main

func maxSatisfied(customers []int, grumpy []int, minutes int) (ret int) {
	pre, post := make([]int, len(customers)+1), make([]int, len(customers)+1)
	for i := 1; i < len(pre); i++ {
		if pre[i] = pre[i-1]; grumpy[i-1] == 0 {
			pre[i] += customers[i-1]
		}
	}
	for i := len(post) - 2; i >= 0; i-- {
		if post[i] = post[i+1]; grumpy[i] == 0 {
			post[i] += customers[i]
		}
	}
	for i := 0; i < minutes; i++ {
		ret += customers[i]
	}
	ret += post[minutes]
	for i, window := minutes, ret-post[minutes]; i < len(customers); i++ {
		window += customers[i] - customers[i-minutes]
		if max := window + pre[i-minutes+1] + post[i+1]; max > ret {
			ret = max
		}
	}
	return ret
}

func main() {}
