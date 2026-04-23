// https://leetcode.com/problems/sum-of-distances/
package main

func distance(nums []int) []int64 {
	res := make([]int64, len(nums))
	cache := make(map[int][]int)
	for i := range nums {
		cache[nums[i]] = append(cache[nums[i]], i)
	}
	tmp := make([]int, 0, len(nums))
	for _, indices := range cache {
		tmp = append(tmp, indices[0])
		for i := 1; i < len(indices); i++ {
			tmp = append(tmp, indices[i]+tmp[i-1])
		}
		for i, index := range indices {
			var left, right int64
			if i > 0 {
				left = int64(i*index - tmp[i-1])
			}
			if n := len(indices); i < len(indices) {
				if right = int64((tmp[n-1] - tmp[i]) - (n-1-i)*index); right < 0 {
					right = -right
				}
			}
			res[index] = left + right
		}
		tmp = tmp[:0]
	}
	return res
}

func main() {}
