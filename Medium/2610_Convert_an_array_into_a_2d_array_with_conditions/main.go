// https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/
package main

func findMatrix(nums []int) [][]int {
	var (
		freq [201]int
		size int
	)
	for _, num := range nums {
		if freq[num]++; freq[num] > size {
			size = freq[num]
		}
	}
	ret := make([][]int, size)
	for num, f := range freq {
		for i := 0; f > 0; i, f = i+1, f-1 {
			ret[i] = append(ret[i], num)
		}
	}
	return ret
}

func main() {}
