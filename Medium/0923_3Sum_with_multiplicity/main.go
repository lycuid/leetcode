// https://leetcode.com/problems/3sum-with-multiplicity/
package main

import "sort"

const MOD = 1e9 + 7

func threeSumMulti(arr []int, target int) (res int) {
	sort.Ints(arr)
	l := len(arr)
	freq := make(map[int]int)
	for i := 0; i < l; i++ {
		freq[arr[i]]++
	}
	for i := 0; i < l; i += freq[arr[i]] {
		j, k := i, l-1
		for j < k {
			if arr[i]+arr[j]+arr[k] < target {
				j += freq[arr[j]]
			} else if arr[i]+arr[j]+arr[k] > target {
				k -= freq[arr[k]]
			} else {
				if arr[i] != arr[j] && arr[j] != arr[k] && arr[i] != arr[k] {
					res += freq[arr[i]] * freq[arr[j]] * freq[arr[k]]
				} else if arr[i] == arr[j] && arr[j] != arr[k] {
					res += (freq[arr[i]] * (freq[arr[i]] - 1) / 2) * freq[arr[k]]
				} else if arr[i] != arr[j] && arr[j] == arr[k] {
					res += (freq[arr[j]] * (freq[arr[j]] - 1) / 2) * freq[arr[i]]
				} else { // arr[i] == arr[j] == arr[k]
					res += freq[arr[i]] * (freq[arr[i]] - 1) * (freq[arr[i]] - 2) / 6
				}
				res %= MOD
				j += freq[arr[j]]
				k -= freq[arr[k]]
			}
		}
	}
	return res % MOD
}

func main() {}
