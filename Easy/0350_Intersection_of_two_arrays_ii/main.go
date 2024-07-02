// https://leetcode.com/problems/intersection-of-two-arrays-ii/
package main

func intersect(nums1 []int, nums2 []int) (ret []int) {
	cache := make(map[int]int)
	for _, num := range nums1 {
		cache[num]++
	}
	for _, num := range nums2 {
		if cache[num] > 0 {
			ret, cache[num] = append(ret, num), cache[num]-1
		}
	}
	return ret
}

func main() {}
