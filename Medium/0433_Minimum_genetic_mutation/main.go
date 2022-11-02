// https://leetcode.com/problems/minimum-genetic-mutation/
package main

type Tuple struct {
	val   string
	stage int
}

func Compare(s1, s2 string) (ret int) {
	for i := range s1 {
		if s1[i] != s2[i] {
			ret++
		}
	}
	return ret
}

func minMutation(start, end string, bank []string) int {
	queue, n := []Tuple{{start, 0}}, len(bank)
	for ; len(queue) > 0; queue = queue[1:] {
		for i, t := 0, queue[0]; i < n; i++ {
			if Compare(t.val, bank[i]) == 1 {
				if bank[i] == end {
					return t.stage + 1
				}
				queue = append(queue, Tuple{bank[i], t.stage + 1})
				bank[i], bank[n-1], n, i = bank[n-1], bank[i], n-1, i-1
			}
		}
	}
	return -1
}

func main() {}
