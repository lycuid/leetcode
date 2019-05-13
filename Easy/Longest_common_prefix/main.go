// https://leetcode.com/problems/longest-common-prefix/

package main

import (
	"bytes"
)

func getTheLongest(arr []string, x int, y int) string {
	var (
		length int
		bytes  bytes.Buffer
	)

	if len(arr[x]) > len(arr[y]) {
		length = len(arr[y])
	} else {
		length = len(arr[x])
	}

	for i := 0; i < length; i++ {
		if arr[x][i] == arr[y][i] {
			bytes.WriteByte(arr[x][i])
		} else {
			break
		}
	}

	return bytes.String()
}

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	substr := strs[0]

	for i := 1; i < length; i++ {
		for j := i - 1; j < length-1; j++ {
			temp := getTheLongest(strs, i, j)
			if len(temp) < len(substr) {
				substr = temp
			}
		}
	}

	return substr

}

func main() {}
