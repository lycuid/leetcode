// https://leetcode.com/problems/string-compression/
package main

import "strconv"

func compress(chars []byte) (i int) {
	for j, k := 0, 1; j < len(chars); j, k = k, k+1 {
		for ; k < len(chars) && chars[k] == chars[j]; k++ {
		}
		chars[i], i = chars[j], i+1
		if n := k - j; n > 1 {
			for _, ch := range []byte(strconv.Itoa(n)) {
				chars[i], i = ch, i+1
			}
		}
	}
	return i
}

func main() {}
