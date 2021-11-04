package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	flag := head
	for i:=0;i<k;i++{
		if flag == nil{
			return head
		}
		flag = flag.Next
	}

	newHead := &ListNode{Next:head}
	nodeList := make([]*ListNode,0)
	//nodeList = append(nodeList,newHead)
	cur := newHead.Next
	prev := newHead
	for cur != nil{
		for i:=0;i<k;i++{

			if cur ==nil{
				return newHead.Next
			}
			nodeList = append(nodeList,cur)
			cur = cur.Next
		}

		nodeList[0].Next = nodeList[len(nodeList)-1].Next
		prev.Next = nodeList[len(nodeList)-1]
		for i:=len(nodeList)-1;i>0;i--{
			fmt.Println(nodeList[i].Val)
			nodeList[i].Next = nodeList[i-1]
		}
		prev = nodeList[0]
		nodeList = []*ListNode{}

	}




	return newHead.Next


}

func main() {
	a:=ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	res := reverseKGroup(&a,2)
	for res != nil{
		fmt.Println(res.Val)
		res = res.Next
	}

}
