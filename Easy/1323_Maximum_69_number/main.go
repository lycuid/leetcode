// https://leetcode.com/problems/maximum-69-number/
package main

func maximum69Number(num int) (ret int) {
	div := 1
	for i := num; i > 0; i /= 10 {
		div *= 10
	}
	for ; div > 0; div /= 10 {
		n := (num / div) * div
		if num, ret = num-n, ret+n; n == 6*div {
			ret += 3 * div
			break
		}
	}
	return ret + num
}

func main() {}
