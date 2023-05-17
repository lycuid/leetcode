// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) (ret int) {
	var (
		fst *ListNode = head
		snd *ListNode = nil
		aux *ListNode = head
	)
	for step := false; head.Next != nil; head = head.Next {
		if step = !step; step {
			snd, aux = aux, aux.Next
		}
	}
	for snd, aux, aux.Next = aux, aux.Next, nil; aux != nil; {
		aux.Next, snd, aux = snd, aux, aux.Next
	}
	for ; snd != nil; fst, snd = fst.Next, snd.Next {
		if value := fst.Val + snd.Val; value > ret {
			ret = value
		}
	}
	return ret
}

func main() {}
