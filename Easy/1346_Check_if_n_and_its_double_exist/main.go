// https://leetcode.com/problems/check-if-n-and-its-double-exist/
package main

func checkIfExist(arr []int) bool {
	cache := make(map[int]bool)
	for _, num := range arr {
		if cache[num*2] {
			return true
		}
		if num%2 == 0 && cache[num/2] {
			return true
		}
		cache[num] = true
	}
	return false
}

func main() {}
