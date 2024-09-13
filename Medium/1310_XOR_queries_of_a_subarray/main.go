// https://leetcode.com/problems/xor-queries-of-a-subarray/
package main

func xorQueries(arr []int, queries [][]int) (ret []int) {
	ret, arr = make([]int, len(queries)), append(arr, 0)
	for i := len(arr) - 2; i >= 0; i-- {
		arr[i] = arr[i] ^ arr[i+1]
	}
	for i, q := range queries {
		ret[i] = arr[q[0]] ^ arr[q[1]+1]
	}
	return ret
}

func main() {}
