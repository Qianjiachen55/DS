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
func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{}
	more := &ListNode{}
	curLess := less
	curMore := more

	cur := head

	for cur != nil{
		if cur.Val <x{
			curLess.Next = cur
			curLess = curLess.Next
		}else{
			curMore.Next = cur
			curMore = curMore.Next
		}
		cur = cur.Next
	}

	curLess.Next = more.Next
	curMore.Next = nil
	return less.Next
}
