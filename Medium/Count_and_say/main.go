// https://leetcode.com/problems/count-and-say/
package main

import "bytes"

func Say(s string) string {
	var res bytes.Buffer
	i, count := 0, '1'
	for i = 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			res.WriteRune(count)
			res.WriteByte(s[i-1])
			count = '1'
			continue
		}
		count++
	}
	res.WriteRune(count)
	res.WriteByte(s[i-1])
	return res.String()
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	return Say(countAndSay(n - 1))
}

func main() {}
