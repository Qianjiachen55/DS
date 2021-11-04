package main

type ListNode struct {
	Val int
	Next *ListNode
}


func swapPairs(head *ListNode) *ListNode {

	if head == nil|| head.Next==nil{
		return head
	}

	newHead := &ListNode{}
	newHead.Next = head
	temp := newHead
	node1 := head
	node2 := head.Next
	for node1 != nil && node2 != nil{
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1

		if temp.Next != nil{
			node1 = node1.Next
			node2 = node1.Next
		}else{
			return newHead.Next
		}

	}
	return newHead.Next
}