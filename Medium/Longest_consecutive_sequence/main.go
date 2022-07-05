// https://leetcode.com/problems/longest-consecutive-sequence/
package main

func longestConsecutive(nums []int) (ret int) {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	for num := range set {
		if !set[num-1] {
			n, l := num, 1
			for set[n+1] {
				n, l = n+1, l+1
			}
			if l > ret {
				ret = l
			}
		}
	}
	return ret
}

func main() {}
