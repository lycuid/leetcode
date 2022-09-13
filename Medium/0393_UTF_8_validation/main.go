// https://leetcode.com/problems/utf-8-validation/
package main

func validUtf8(data []int) bool {
	n, j := len(data), 0
	for i := 0; i < n; {
		if data[i]>>7 == 0 {
			j = i + 1
		} else if data[i]>>5 == 0b110 {
			j = i + 2
		} else if data[i]>>4 == 0b1110 {
			j = i + 3
		} else if data[i]>>3 == 0b11110 {
			j = i + 4
		} else {
			return false
		}
		if j > n {
			return false
		}
		for i++; i < j; i++ {
			if data[i]>>6 != 0b10 {
				return false
			}
		}
	}
	return true
}

func main() { }
