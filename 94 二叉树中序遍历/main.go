package main



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	ans := []int{}

	var inorder func(root *TreeNode)

	inorder = func(root *TreeNode){
		if root == nil{
			return
		}

		//left
		inorder(root.Left)
		// root
		ans = append(ans,root.Val)
		//right
		inorder(root.Right)
	}

	inorder(root)


	return ans

}
