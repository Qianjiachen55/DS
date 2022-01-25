package main

import (
	"container/list"
	"math"
)

type MinStack struct {
	stack *list.List
	minStack *list.List
}

func Constructor() MinStack {
	m := MinStack{
		stack: list.New(),
		minStack: list.New(),
	}
	m.minStack.PushBack(math.MaxInt64)
	return m
}


func (this *MinStack) Push(x int)  {
	this.stack.PushBack(x)
	if this.minStack.Back().Value.(int) >=x{
		this.minStack.PushBack(x)
	}
}


func (this *MinStack) Pop()  {
	x := this.stack.Back()
	this.stack.Remove(x)
	if x.Value.(int)==this.minStack.Back().Value.(int){
		temp := this.minStack.Back()
		this.minStack.Remove(temp)
	}

}


func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}


func (this *MinStack) Min() int {
	return this.minStack.Back().Value.(int)
}
