// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/
package main

import (
	"fmt"
	"strconv"
)

func getLucky(s string, k int) int {
	var str string
	for _, ch := range s {
		str += fmt.Sprintf("%d", int(ch-'a'+1))
	}
	for ; k > 0; k-- {
		var num int
		for _, ch := range str {
			num += int(ch - '0')
		}
		str = fmt.Sprintf("%d", num)
	}
	num, _ := strconv.Atoi(str)
	return num
}

func main() {}
