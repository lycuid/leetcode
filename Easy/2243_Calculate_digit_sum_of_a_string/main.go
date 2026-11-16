// https://leetcode.com/problems/calculate-digit-sum-of-a-string/
package main

func digitSum(s string, k int) string {
	curr, prev := make([]byte, 0, len(s)), make([]byte, 0, len(s))
	for i := range s {
		curr = append(curr, s[i])
	}
	cache := make([]byte, 0, k)
	for len(curr) > k {
		prev, curr = curr, prev
		curr = curr[:0]
		for i := 0; i < len(prev); i += k {
			var tmp int
			for j := 0; j < k && i+j < len(prev); j++ {
				tmp += int(prev[i+j] - '0')
			}
			cache = cache[:0]
			for {
				cache = append(cache, byte(tmp%10)+'0')
				if tmp /= 10; tmp == 0 {
					break
				}
			}
			for l := len(cache) - 1; l >= 0; l-- {
				curr = append(curr, cache[l])
			}
		}
	}
	return string(curr)
}

func main() {}
