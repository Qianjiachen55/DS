package main


type LRUCache struct {
	capacity int
	head,tail *Node
	m map[int]*Node
	curr int
}


type Node struct {
	prev *Node
	next *Node
	key int
	Val int
}

func Constructor(capacity int) LRUCache {

	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	m := map[int]*Node{}

	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		m:        m,
		curr:     0,
	}
}


func (this *LRUCache) Get(key int) int {
	if _,ok := this.m[key];ok{
		this.moveHead(key)
		return this.m[key].Val
	}

	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok := this.m[key];ok{
		// update
		this.m[key].Val = value
		this.moveHead(key)
		return
	}

	if this.curr < this.capacity{
		//addToHead
		this.addHead(key,value)
		this.curr++
		return
	}else{
		//addToHead && deleteTail
		this.addHead(key,value)
		this.deleteTail()
		return
	}
}


func (this *LRUCache) addHead(key int,value int){
	node := &Node{key: key,Val: value}
	node.next = this.head.next
	node.prev = this.head
	node.next.prev = node
	this.head.next = node
	this.m[key] = node
}

func (this *LRUCache) moveHead (key int){
	//delete && add
	val := this.m[key].Val
	this.delete(key)
	this.addHead(key,val)

}

func (this *LRUCache) deleteTail() {
	node := this.tail.prev
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(this.m,node.key)
}

func (this *LRUCache) delete(key int) {
	node := this.m[key]
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(this.m,node.key)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */