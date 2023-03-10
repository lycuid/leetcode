// https://leetcode.com/problems/linked-list-random-node/
package main

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct{ inner []int }

func Constructor(head *ListNode) Solution {
	var inner []int
	for ; head != nil; head = head.Next {
		inner = append(inner, head.Val)
	}
	return Solution{inner}
}

func (this *Solution) GetRandom() int {
	return this.inner[rand.Intn(len(this.inner))]
}

func main() {}
