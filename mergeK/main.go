package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for i := 0; i < len(lists); i++ {
		cur.Next = mergeTwoList(newHead.Next ,lists[i])
	}

	return newHead.Next

}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			return newHead.Next
		} else if l2 == nil {
			cur.Next = l1
			return newHead.Next
		} else if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		}

	}

	return newHead.Next

}

func main()  {
	temp := []*ListNode{
		&ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	res :=mergeKLists(temp)
	for res != nil{
		fmt.Println(res.Val)
		res = res.Next
	}
}
