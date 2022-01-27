package main

type LRUCache struct {
	head,tail *Node
	size int
	capacity int
	cache map[int]*Node
}


type Node struct {
	prev,next *Node
	key,value int
}

func initNode(key, value int) *Node  {
	return &Node{
		key:   key,
		value: value,
	}
}
func Constructor(capacity int)LRUCache  {
	l := LRUCache{
		head:     initNode(0,0),
		tail:     initNode(0,0),
		size:     0,
		capacity: capacity,
		cache:    map[int]*Node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}
func (this *LRUCache) Get(key int)int {
	if _,ok := this.cache[key];!ok{
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache)Put(key ,value int)  {
	if _,ok := this.cache[key];!ok{
		//add
		node := initNode(key,value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity{
			removeNode := this.removeTail()
			delete(this.cache,removeNode.key)
			this.size--
		}
	}else{
		//update

		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}

}

/* put
1. update
	removeNode&&addToHead
2. add
	addToHead(&& removeTail)
*/

/* Get
moveToHead
	removeNode&&addToHead
*/
func (this *LRUCache)moveToHead(node *Node)  {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *Node)  {
	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node

}

func (this *LRUCache) removeTail() *Node {
	node:=this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

}
