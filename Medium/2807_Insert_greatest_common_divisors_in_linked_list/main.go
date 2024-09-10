// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Aux(node, prev *ListNode) *ListNode {
	if node != nil {
		if node.Next = Aux(node.Next, node); prev != nil {
			return &ListNode{
				Val:  gcd(prev.Val, node.Val),
				Next: node,
			}
		}
	}
	return node
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	return Aux(head, nil)
}

func main() {}
