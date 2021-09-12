package list

type Node struct {
	value int
	next  *Node
}

var Head = Node{value: 23, next: nil}
var Tail = Node{}

func (n *Node) Prepend(value int) *Node {

	return &Node{}
}

func (n *Node) Append(value int) *Node {
	return &Node{}
}
