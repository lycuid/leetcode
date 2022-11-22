// https://leetcode.com/problems/perfect-squares/
package main

const UNREACHABLE, UNVISITED int = 1e5 + 7, -1

func Find(n int, squares []int, cache []int) int {
	if l := 0; cache[n] == UNVISITED {
		for r := len(squares) - 1; l < r; {
			if mid := (l + r) / 2; squares[mid] <= n {
				l = mid + 1
			} else if squares[mid] > n {
				r = mid
			}
		}
		for cache[n] = UNREACHABLE; l >= 0; l-- {
			if m := n - squares[l]; m >= 0 {
				if cache[m] = Find(m, squares, cache); cache[m]+1 < cache[n] {
					cache[n] = cache[m] + 1
				}
			}
		}
	}
	return cache[n]
}

func numSquares(n int) int {
	cache := make([]int, n+1)
	for i := 1; i < len(cache); i++ {
		cache[i] = UNVISITED
	}
	var squares []int
	for i := 1; i*i <= n; i++ {
		squares, cache[i*i] = append(squares, i*i), 1
	}
	return Find(n, squares, cache)
}

func main() {}
