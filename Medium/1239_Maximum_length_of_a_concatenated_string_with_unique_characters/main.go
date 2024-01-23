// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
package main

func Aux(arr []string, cache []uint32, value uint32) (max int) {
	for i := range cache {
		if cache[i]>>27&1 == 0 && value&cache[i] == 0 {
			if n := len(arr[i]) + Aux(arr[i+1:], cache[i+1:], value|cache[i]); n > max {
				max = n
			}
		}
	}
	return max
}

func maxLength(arr []string) int {
	cache := make([]uint32, len(arr))
	for i := range arr {
		for j := range arr[i] {
			var n uint32 = 1 << (arr[i][j] - 'a')
			if cache[i]&n != 0 {
				n = 1 << 27
			}
			cache[i] |= n
		}
	}
	return Aux(arr, cache, 0)
}

func main() {}
