// https://leetcode.com/problems/fibonacci-number/
package main

func fib(n int) int {
	m_1, m := 0, 1
	for ; n > 0; n-- {
		m_1, m = m, m+m_1
	}
	return m_1
}
