package main



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func isSameTree(p *TreeNode, q *TreeNode) bool {
	return helper(p,q)

}

func helper(p * TreeNode,q *TreeNode)bool{
	if p == nil || q == nil{
		if p ==nil && q== nil{
			return true
		}
		return false
	}

	if p.Val == q.Val{
		return helper(p.Right,q.Right) && helper(p.Left,q.Left)
	}

	return false
}
