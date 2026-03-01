// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
package main

func minPartitions(n string) (num int) {
	for i := range n {
		num = max(num, int(n[i]-'0'))
	}
	return num
}

func main() {}
