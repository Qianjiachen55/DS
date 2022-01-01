package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}

	return helper(root.Left,root.Right)
}

func helper(p *TreeNode, q *TreeNode) bool{
	if p == nil || q == nil{
		if p== nil && q == nil{
			return true
		}
		return false
	}

	if p.Val == q.Val{
		return helper(p.Left,q.Right) && helper(q.Left,p.Right)
	}

	return false


}
