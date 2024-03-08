// https://leetcode.com/problems/count-elements-with-maximum-frequency/
package main

func maxFrequencyElements(nums []int) (ret int) {
	freq, max := make(map[int]int), 0
	for _, num := range nums {
		if freq[num]++; freq[num] > max {
			max = freq[num]
		}
	}
	for _, f := range freq {
		if f == max {
			ret += f
		}
	}
	return ret
}

func main() {}
