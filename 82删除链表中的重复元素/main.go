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
		Val:-1000,
		Next:head,
	}
	cur := newHead
	temp := 0
	for cur.Next!=nil && cur.Next.Next !=nil{
		if cur.Next.Val ==cur.Next.Next.Val{
			temp = cur.Next.Val
			for cur.Next != nil && cur.Next.Val == temp{
				cur.Next = cur.Next.Next
			}
		}else{
			cur = cur.Next
		}
	}



	return newHead.Next


}
