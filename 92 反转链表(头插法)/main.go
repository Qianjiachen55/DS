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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val:-1}
	dummyNode.Next = head
	pre :=dummyNode

	count := 1
	for count<left{
		pre = pre.Next
		count ++
	}
	cur := pre.Next

	for count < right{
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

		count++
	}

	return dummyNode.Next
}
