// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) (num int) {
	for ; head != nil; head = head.Next {
		num = (num << 1) + head.Val
	}
	return num
}

func main() {}
