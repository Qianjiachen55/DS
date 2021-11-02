package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next:head}
	cur := newHead
	pre :=cur
	for cur.Next != nil{
		if n>0{
			n--
		}else{
			pre = pre.Next
		}
		cur = cur.Next
	}
	pre.Next = pre.Next.Next

	return newHead.Next

}
