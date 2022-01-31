package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	a1 := &ListNode{
		Val:  4,
		Next: nil,
	}
	a2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	a3 := &ListNode{
		Val:  1,
		Next: nil,
	}
	a4 := &ListNode{
		Val:  3,
		Next: nil,
	}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4

	ans := sortList(a1)

	for ans != nil{
		fmt.Println(ans.Val)
		ans  = ans.Next
	}
}


func sortList(head *ListNode) *ListNode {
	arr := []*ListNode{}
	for head != nil{
		arr = append(arr,head)
		head = head.Next
	}
	if len(arr)==0{
		return nil
	}
	Qsort(0,len(arr)-1,arr)
	newHead := arr[0]
	for i:=1;i<len(arr)-1;i++{
		arr[i].Next = arr[i+1]

	}
	arr[len(arr)-1] = nil

	//arr[len(arr)-1] = nil

	return newHead

}


func Qsort(start,end int,arr []*ListNode){
	if start >= end{
		return
	}

	index := partition(start,end,arr)
	Qsort(start,index-1,arr)
	Qsort(index+1,end,arr)
}

func partition(start,end int,arr []*ListNode)int{
	temp := arr[start]
	left,rigth := start,end
	for left < rigth{
		for left < rigth && arr[rigth].Val>=temp.Val{
			rigth --
		}
		for left <rigth && arr[left].Val<=temp.Val{
			left ++
		}

		arr[left],arr[rigth] = arr[rigth],arr[left]
	}

	arr[start],arr[left] = arr[left],arr[start]

	return left
}
