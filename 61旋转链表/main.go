package main



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func rotateRight(head *ListNode, k int) *ListNode {
	if head ==nil||head.Next==nil||k==0{
		return head
	}
	n:=1
	cur := &ListNode{
		Next :head.Next,
	}
	for cur.Next!=nil{
		cur = cur.Next
		n++
	}
	cur.Next = head

	k = n-k%n
	res := head
	for i:=1;i<k;i++{
		res  = res.Next
	}
	res.Next,res = nil,res.Next

	return res

}
