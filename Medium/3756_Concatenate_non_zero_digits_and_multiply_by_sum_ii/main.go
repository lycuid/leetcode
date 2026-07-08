// https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-ii/
package main

func sumAndMultiply(s string, queries [][]int) []int {
	const mod = 1e9 + 7
	var (
		res    = make([]int, len(queries))
		prefix = make([]int, len(s)+1)
		digits = make([]int, 1, len(s)+1)
		count  = make([]int, len(s)+1)
	)
	for i, ch := range s {
		prefix[i+1], count[i+1] = prefix[i+1]+prefix[i], count[i]
		if num := int(ch - '0'); num > 0 {
			prefix[i+1] = (prefix[i+1] + num) % mod
			count[i+1]++
			digits = append(digits, (digits[len(digits)-1]*10+num)%mod)
		}
	}
	pow10 := make([]int, len(digits))
	pow10[0] = 1
	for i := 1; i < len(pow10); i++ {
		pow10[i] = (pow10[i-1] * 10) % mod
	}
	for i, query := range queries {
		drop, take := count[query[0]], count[query[1]+1]
		num := (digits[take] - (digits[drop] * pow10[take-drop])) % mod
		if num < 0 {
			num += mod
		}
		s := prefix[query[1]+1] - prefix[query[0]]
		res[i] = (num * s) % mod
	}
	return res
}

func main() {}
