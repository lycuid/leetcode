// https://leetcode.com/problems/rotate-image/
package main

func rotate(m [][]int) {
	s, e := 0, len(m)-1
	for e >= s {
		for i := 0; i < e-s; i++ {
			temp := m[s][s+i]
			m[s][s+i] = m[e-i][s]
			m[e-i][s] = m[e][e-i]
			m[e][e-i] = m[s+i][e]
			m[s+i][e] = temp
		}
		s++
		e--
	}
}

func main() {}
