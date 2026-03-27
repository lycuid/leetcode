// https://leetcode.com/problems/matrix-similarity-after-cyclic-shifts/
package main

func areSimilar(mat [][]int, k int) bool {
	if len(mat) > 0 {
		n := len(mat[0])
		if k = k % n; k > 0 {
			equal := func(m1, m2 []int) bool {
				for i := range m1 {
					if m1[i] != m2[i] {
						return false
					}
				}
				return true
			}

			cache := make([]int, n)

			reverse := func(m []int, k int) []int {
				for i := range m {
					if i < k {
						cache[k-1-i] = m[i]
					} else {
						cache[k+n-1-i] = m[i]
					}
				}
				for i := 0; i < n/2; i++ {
					cache[i], cache[n-1-i] = cache[n-1-i], cache[i]
				}
				return cache
			}

			for i := range mat {
				if _k := k % n; _k > 0 {
					if i%2 != 0 {
						_k = n - _k
					}
					if !equal(mat[i], reverse(mat[i], _k)) {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {}
