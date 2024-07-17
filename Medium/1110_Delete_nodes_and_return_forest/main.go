// https://leetcode.com/problems/delete-nodes-and-return-forest/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) (ret []*TreeNode) {
	root = &TreeNode{Val: 0, Left: root}
	cache := map[int]bool{0: true}
	for _, val := range to_delete {
		cache[val] = true
	}

	var Aux func(*TreeNode) *TreeNode
	Aux = func(root *TreeNode) *TreeNode {
		if root != nil {
			root.Left = Aux(root.Left)
			root.Right = Aux(root.Right)
			if cache[root.Val] {
				if root.Left != nil {
					ret = append(ret, root.Left)
				}
				if root.Right != nil {
					ret = append(ret, root.Right)
				}
				return nil
			}
		}
		return root
	}
	Aux(root)

	return ret
}

func main() {}
