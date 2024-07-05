// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	ret := []int{math.MaxInt, 0}
	if head != nil {
		node, prev, f_index, index := head.Next, head, -1, -1
		for i := 1; node != nil; node, prev, i = node.Next, node, i+1 {
			if next := node.Next; next != nil {
				if (node.Val > prev.Val && node.Val > next.Val) ||
					(node.Val < prev.Val && node.Val < next.Val) {
					if index != -1 {
						ret[0] = min(ret[0], i-index)
					} else {
						f_index = i
					}
					index = i
				}
			}
		}
		ret[1] = index - f_index
	}
	if ret[0] == math.MaxInt {
		ret[0] = -1
	}
	if ret[1] == 0 {
		ret[1] = -1
	}
	return ret
}

func main() {}
