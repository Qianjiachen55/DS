package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode{root}
	target := false

	if root == nil{
		return res
	}

	for len(queue) > 0{
		resTemp := []int{}
		size := len(queue)
		for i:=0;i<size;i++{
			// if target{
			//     resTemp = append([]int{queue[i].Val},resTemp...)
			// }else{
			//     resTemp = append(resTemp,queue[i].Val)
			// }
			resTemp = append(resTemp,queue[i].Val)

			if queue[i].Left != nil{
				queue = append(queue,queue[i].Left)
			}
			if queue[i].Right != nil{
				queue = append(queue,queue[i].Right)
			}
		}
		if target{
			for j := 0;j<size/2;j++{
				resTemp[j],resTemp[size-j-1] = resTemp[size-j-1],resTemp[j]
			}
		}


		res = append(res,resTemp)
		target = !target
		queue=queue[size:]
	}

	return res
}
