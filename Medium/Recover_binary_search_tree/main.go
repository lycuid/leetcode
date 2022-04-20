// https://leetcode.com/problems/recover-binary-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeSolver struct {
	prev, fst, snd *TreeNode
}

func (s *TreeSolver) inorder(root *TreeNode) {
	if root == nil {
		return
	}
	s.inorder(root.Left)
	if s.prev != nil && root.Val < s.prev.Val {
		if s.fst == nil {
			s.fst = s.prev
		}
		s.snd = root
	}
	s.prev = root
	s.inorder(root.Right)
}

func recoverTree(root *TreeNode) {
	var solver TreeSolver
	solver.inorder(root)
	if solver.fst != nil && solver.snd != nil {
		solver.fst.Val, solver.snd.Val = solver.snd.Val, solver.fst.Val
	}
}

func main() {}
