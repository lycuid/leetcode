// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
package main

func minPartitions(n string) int {
	ret := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > ret {
			ret = n[i]
		}
	}
	return int(ret - '0')
}

func main() {}
