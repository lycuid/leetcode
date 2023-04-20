// https://leetcode.com/problems/maximum-width-of-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct{ inner []*TreeNode }

func (q Queue) Empty() bool          { return len(q.inner) == 0 }
func (q *Queue) Push(item *TreeNode) { q.inner = append(q.inner, item) }
func (q *Queue) Pop() (item *TreeNode) {
	if !q.Empty() {
		item, q.inner = q.inner[0], q.inner[1:]
	}
	return item
}

func widthOfBinaryTree(root *TreeNode) (ret int) {
	var queue [2]Queue
	var l, r int
	root.Val = 1
	for queue[0].Push(root); !queue[0].Empty(); queue[0], queue[1] = queue[1], queue[0] {
		node := queue[0].Pop()
		for l, r = node.Val, node.Val; node != nil; r, node = node.Val, queue[0].Pop() {
			if left := node.Left; left != nil {
				left.Val = 2 * node.Val
				queue[1].Push(left)
			}
			if right := node.Right; right != nil {
				right.Val = 2*node.Val + 1
				queue[1].Push(right)
			}
		}
		if n := r - l + 1; n > ret {
			ret = n
		}
	}
	return ret
}

func main() {}
