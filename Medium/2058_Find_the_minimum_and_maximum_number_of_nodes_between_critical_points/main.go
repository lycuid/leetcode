// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	res := []int{-1, -1}
	if head != nil {
		var fstId, prevId, currId int
		prev, curr := head, head.Next
		for id := 1; curr != nil; prev, curr, id = curr, curr.Next, id+1 {
			if next := curr.Next; next != nil && ((prev.Val < curr.Val && next.Val < curr.Val) || (prev.Val > curr.Val && next.Val > curr.Val)) {
				if fstId == 0 {
					fstId = id
				}
				if currId != 0 {
					prevId = currId
				}
				currId = id
				if prevId != 0 {
					if res[0] == -1 {
						res[0] = currId - prevId
					} else {
						res[0] = min(res[0], currId-prevId)
					}
				}
			}
		}
		if fstId != currId {
			res[1] = currId - fstId
		}
	}
	return res
}

func main() {}
