// https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/
package main

func numOfSubarrays(arr []int) (res int) {
	var (
		prev, curr int
		count      = [2]int{1, 0}
	)
	for _, num := range arr {
		prev, curr = curr, prev
		curr = prev + num
		count[curr&1]++
		res = (res + count[1-curr&1]) % (1e9 + 7)
	}
	return res
}

func main() {}
