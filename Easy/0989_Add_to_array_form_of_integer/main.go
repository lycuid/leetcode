// https://leetcode.com/problems/add-to-array-form-of-integer/
package main

func addToArrayForm(num []int, k int) (ret []int) {
	var carry int
	for i := len(num) - 1; i >= 0; i, k = i-1, k/10 {
		n := num[i] + k%10 + carry
		ret, carry = append(ret, n%10), n/10
	}
	for ; k > 0 || carry > 0; k = k / 10 {
		n := k%10 + carry
		ret, carry = append(ret, n%10), n/10
	}
	for i, n := 0, len(ret); i < n/2; i++ {
		ret[i], ret[n-1-i] = ret[n-1-i], ret[i]
	}
	return ret
}

func main() {}
