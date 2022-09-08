// https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
package main

func hasAllCodes(s string, k int) bool {
	max_subs := 1 << k
	if len(s) < max_subs+k-1 {
		return false
	}
	set, num := make([]bool, max_subs+1), 0
	for i := 0; i < k; i++ {
		num = ((num << 1) + int(s[i]-'0')) & ((1 << k) - 1)
	}
	set[num], max_subs = true, max_subs-1
	for i := k; i < len(s); i++ {
		num = ((num << 1) + int(s[i]-'0')) & ((1 << k) - 1)
		if !set[num] {
			set[num], max_subs = true, max_subs-1
		}
	}
	return max_subs == 0
}

func main() {}
