// https://leetcode.com/problems/minimum-absolute-distance-between-mirror-pairs/
package main

func minMirrorPairDistance(nums []int) int {
	var (
		res     = len(nums)
		cache   = make(map[int]int)
		reverse = func(num int) (res int) {
			for ; num > 0; num /= 10 {
				res = res*10 + num%10
			}
			return res
		}
	)
	for i, num := range nums {
		if index, found := cache[num]; found {
			res = min(res, i-index)
		}
		cache[reverse(num)] = i
	}
	if res == len(nums) {
		return -1
	}
	return res
}

func main() {}
