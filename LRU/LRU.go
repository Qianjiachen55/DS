package main

import "fmt"

type Node struct {
	Key   int
	Value int
	Next  *Node
	Pre   *Node
}

type LRUMap struct {
	head *Node
	tail *Node
	m    map[int]*Node
	size int
}

func (p *LRUMap) Set(key int, value int) {
	if node, ok := p.m[key]; ok {
		//此key存在
		node.Value = value
		if node == p.head {
			return
		} else {
			p.SetHead(node)
			//将此节点换到头节点
		}
	} else {
		//key不存在

		if len(p.m) == p.size {
			//长度处理
			//已满，去除一个
			tailKey := p.tail.Key
			tailNew := p.tail.Pre
			fmt.Println("tailNew",tailNew)
			fmt.Println(p.tail.Key,p.tail.Pre)
			delete(p.m, tailKey)
			tailNew.Next = nil
			p.tail = tailNew
		}

		node := Node{
			Key:   key,
			Value: value,
		}
		//将node设置为head

		if p.head == nil {
			//头节点不存在
			p.head = &node
			p.tail = &node

		} else {
			//头节点存在
			p.SetHead(&node)
		}

		p.m[key] = &node
	}

}

func (p *LRUMap) SetHead(node *Node) {
	//head 为 nil
	if node == p.head{
		return
	}
	if node == p.tail{
		prev := p.tail.Pre
		prev.Next = nil
		node.Pre = nil
		node.Next = p.head
		p.head.Pre = node
		p.head = node

	}
	pre := node.Pre
	next := node.Next
	node.Next = nil
	node.Pre = nil

	if pre != nil {
		pre.Next = next
	}
	if next != nil {
		next.Pre = pre
	}

	node.Next = p.head
	p.head.Pre = node
	p.head = node




}

func (p *LRUMap) Get(key int) int {
	res := -1
	if node, ok := p.m[key]; ok {
		res = node.Value
		//拉到头节点
		if node != p.head {
			p.SetHead(node)
		}
	}
	return res

}

//func (*LRUMap)init(k int) *LRUMap {
//	return &LRUMap{
//		head: nil,
//		tail: nil,
//		m:    make(map[int]*Node,k),
//		size: k,
//	}
//}

func LRU(operators [][]int, k int) []int {
	// write code here
	lru := LRUMap{
		head: nil,
		tail: nil,
		m:    make(map[int]*Node, k),
		size: k,
	}
	res := make([]int,0)
	fmt.Println()
	for _,operator := range operators{
		if operator[0] == 1 {
			fmt.Println("op",1)
			lru.Set(operator[1], operator[2])
		}else if operator[0] ==2 {
			fmt.Println("op",2)
			a :=lru.Get(operator[1])
			res = append(res, a)
		}
		fmt.Println(lru.m)
	}
	//print(res)
	return res
}

func main() {
	var slice [][]int
	slice = [][]int{{1,1,1},{1,2,2},{1,3,2},{2,1},{1,4,4},{2,2}}
	print(slice)
	res := LRU(slice,3)
	fmt.Println(res)
}