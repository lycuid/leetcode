// https://leetcode.com/problems/minimum-penalty-for-a-shop/
package main

func bestClosingTime(customers string) (ret int) {
	post := make([]int, len(customers)+1)
	for i := len(customers) - 1; i >= 0; i-- {
		if post[i] = post[i+1]; customers[i] == 'Y' {
			post[i]++
		}
	}
	pre, penalty := 0, post[0]
	for i := 1; i < len(post); i++ {
		if customers[i-1] == 'N' {
			pre++
		}
		if n := pre + post[i]; n < penalty {
			penalty, ret = n, i
		}
	}
	return ret
}

func main() {}
