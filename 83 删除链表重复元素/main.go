package main


type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head ==nil{
		return head
	}
	newHead := &ListNode{
		Next:head,
	}

	cur := newHead.Next

	for cur.Next !=nil{
		if cur.Val==cur.Next.Val{
			cur.Next = cur.Next.Next
		}else{
			cur= cur.Next
		}
	}

	return newHead.Next
}
