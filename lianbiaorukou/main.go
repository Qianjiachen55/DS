package main

import "fmt"

type ListNode struct {
	value int
	next *ListNode
}

func main() {
	head := ListNode{
		value: 3,
	}
	node2 := ListNode{
		value: 2,
		next:  nil,
	}
	node3 := ListNode{
		value: 0,
		next:  nil,
	}
	node4 := ListNode{
		value: 4,
		next: nil,
	}
	head.next = &node2
	node2.next = &node3
	node3.next = &node4
	node4.next = &node2

	node := Demo_hash(&head)
	fmt.Println(node)
	fmt.Println(Demo_fast_slow(&head))
}


func Demo_fast_slow(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil{
		slow = slow.next
		if fast.next == nil{
			return nil
		}
		fast = fast.next.next
		if slow== fast{
			p := head
			for p!= slow{
				p = p.next
				slow = slow.next
			}
			return p
		}
	}
	return nil
}

func Demo_hash(head *ListNode) *ListNode {
	ListNodeMap := make(map[*ListNode]struct{}, 0)
	node := head
	for node != nil {
		if _, ok := ListNodeMap[node]; ok {
			return node
		}
		ListNodeMap[node] = struct{}{}
		node = node.next
	}
	return nil
}
