// https://leetcode.com/problems/remove-k-digits/
package main

func removeKdigits(num string, k int) string {
	ns := []byte{'0'}
	for i := range num {
		for ; k > 0 && len(ns) > 0 && ns[len(ns)-1] > num[i]; k-- {
			ns = ns[:len(ns)-1]
		}
		ns = append(ns, num[i])
	}
	for ; k > 0; k-- {
		ns = ns[:len(ns)-1]
	}
	for len(ns) > 1 && ns[0] == '0' {
		ns = ns[1:]
	}
	return string(ns)
}

func main() {}
