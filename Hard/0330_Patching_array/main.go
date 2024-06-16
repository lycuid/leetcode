// https://leetcode.com/problems/patching-array/
package main

func minPatches(nums []int, n int) (patches int) {
	var sum int
	for _, num := range nums {
		if num > n {
			break
		}
		for ; sum < num-1; sum += sum + 1 {
			patches++
		}
		sum += num
	}
	for ; sum < n; sum += sum + 1 {
		patches++
	}
	return patches
}

func main() {}
