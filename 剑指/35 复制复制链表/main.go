package main


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)

	for cur := head;cur != nil;cur = cur.Next{
		m[cur] = &Node{Val:cur.Val}
	}

	for cur:=head;cur !=nil;cur = cur.Next{
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
	}

	return m[head]
}
