// https://leetcode.com/problems/single-number-iii/
package main

func singleNumber(nums []int) (ret []int) {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	for num, count := range freq {
		if count == 1 {
			ret = append(ret, num)
		}
	}
	return ret
}

func main() {}
