package main

import (
	"fmt"

	linkedlist "github.com/sktandon0121/go-learning/linkedlist"
)

func main() {
	fmt.Println("Hello Go")
	newList := linkedlist.NewLinkedList()
	node4 := linkedlist.Node{Value: 19}
	newList.Prepend(&node4)
	node3 := linkedlist.Node{Value: 12}
	newList.Prepend(&node3)

	node1 := linkedlist.Node{Value: 30}
	node2 := linkedlist.Node{Value: 17}
	node5 := linkedlist.Node{Value: 10}
	newList.Append(&node1)
	newList.Append(&node2)
	newList.Append(&node5)

	node6 := linkedlist.Node{Value: 8}
	newList.Prepend(&node6)

	// Insert at index
	node7 := linkedlist.Node{Value: 90}
	node8 := linkedlist.Node{Value: 100}
	newList.InsertAt(3, &node7)
	newList.InsertAt(30, &node8)

	newList.TraverseList()

	newList.ReverseList()

	// newList.TraverseList()
	newList.PrintList()
}
