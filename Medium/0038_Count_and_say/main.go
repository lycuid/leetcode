// https://leetcode.com/problems/count-and-say/
package main

import "fmt"

func Slice(str string) (count int, value byte) {
	if len(str) > 0 {
		value = str[0]
		for i := 0; i < len(str) && str[i] == value; i++ {
			count++
		}
	}
	return count, value
}

func countAndSay(n int) (ret string) {
	if n == 1 {
		return "1"
	}
	for prev := countAndSay(n - 1); len(prev) > 0; {
		count, value := Slice(prev)
		ret += fmt.Sprintf("%d%c", count, value)
		prev = prev[count:]
	}
	return ret
}

func main() {}
