// https://leetcode.com/problems/combinations/
package main

func combine(n int, k int) (ret [][]int) {
	if n == 0 || n < k {
		return nil
	}
	if k == 1 {
		ret = append(ret, []int{n})
	}
	for _, combination := range combine(n-1, k-1) {
		ret = append(ret, append([]int{n}, combination...))
	}
	ret = append(ret, combine(n-1, k)...)
	return ret
}

func main() {}
