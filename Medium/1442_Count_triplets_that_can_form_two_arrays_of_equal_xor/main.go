// https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
package main

func countTriplets(arr []int) (count int) {
	for i, xor := range arr {
		for j := i + 1; j < len(arr); j++ {
			if xor ^= arr[j]; xor == 0 {
				count += j - i
			}
		}
	}
	return count
}

func main() {}
