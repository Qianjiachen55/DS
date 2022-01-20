package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil{
		return ret
	}

	queue := []*TreeNode{root}

	for i:=0;len(queue)>0;i++{
		temp := []int{}
		size := len(queue)

		for j:=0;j<size;j++{
			temp = append(temp,queue[j].Val)
			if queue[j].Left != nil{
				queue = append(queue,queue[j].Left)
			}
			if queue[j].Right != nil{
				queue = append(queue,queue[j].Right)
			}
		}

		queue = queue[size:]
		ret = append(ret,temp)
	}

	return ret

}