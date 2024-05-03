// https://leetcode.com/problems/compare-version-numbers/
package main

import (
	"strconv"
	"strings"
)

func Aux(s string) (nums []int) {
	list := strings.Split(s, ".")
	nums = make([]int, len(list))
	for i, num := range list {
		nums[i], _ = strconv.Atoi(num)
	}
	return nums
}

func compareVersion(version1 string, version2 string) int {
	v1, v2 := Aux(version1), Aux(version2)
	for ; len(v1) > 0 && len(v2) > 0; v1, v2 = v1[1:], v2[1:] {
		if a, b := v1[0], v2[0]; a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	for len(v1) > 0 && v1[0] == 0 {
		v1 = v1[1:]
	}
	for len(v2) > 0 && v2[0] == 0 {
		v2 = v2[1:]
	}
	if len(v1) > len(v2) {
		return 1
	}
	if len(v1) < len(v2) {
		return -1
	}
	return 0
}

func main() {}
