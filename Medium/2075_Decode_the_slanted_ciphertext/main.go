// https://leetcode.com/problems/decode-the-slanted-ciphertext/
package main

import "strings"

func decodeCiphertext(encodedText string, rows int) string {
	ret := make([]byte, 0, len(encodedText))
	cols := len(encodedText) / rows
	for i := 0; i < cols; i++ {
		for j, k := 0, i; j < rows && k < cols; j, k = j+1, k+1 {
			ret = append(ret, encodedText[j*cols+k])
		}
	}
	return strings.TrimRight(string(ret), " ")
}

func main() {}
