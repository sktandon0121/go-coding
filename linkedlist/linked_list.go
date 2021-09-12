package linkedlist

import (
	"encoding/json"
	"fmt"
	"log"
)

type LinkedListService interface {
	Append(node *Node)
	Prepend(newNode *Node)
	PrintList()
	TraverseList()
	InsertAt(index int, newNode *Node)
	ReverseList()
}

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList() LinkedListService {
	return &LinkedList{}
}

func (l *LinkedList) Append(node *Node) {

	if l.length == 0 {
		l.head = node
		l.length++
	} else {
		currentNode := l.head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = node
		l.length++
	}
}

func (l *LinkedList) Prepend(newNode *Node) {
	if l.length == 0 {
		l.head = newNode
		l.length++
	} else {
		newNode.Next = l.head
		l.head = newNode
		l.length++
	}
}

func (l *LinkedList) PrintList() {
	data := l
	json, err := json.MarshalIndent(data.head, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
	fmt.Println("lenght of the list is ", data.length)
}

func (l *LinkedList) TraverseList() {
	head := l.head
	list := make([]int, 0, l.length)
	for head != nil {
		list = append(list, head.Value)
		head = head.Next
	}

	fmt.Println(list)
}

func (l *LinkedList) InsertAt(index int, newNode *Node) {
	if l.length == 0 {
		l.head = newNode
		l.length++
	} else if l.length < index {
		l.Append(newNode)
	} else {
		currentNode := l.head
		i := 1
		for i < index-1 {
			currentNode = currentNode.Next
			i++
		}
		nextNode := currentNode.Next
		currentNode.Next = newNode
		newNode.Next = nextNode
		l.length++
	}
}

func (l *LinkedList) ReverseList() {
	pCurr := l.head
	var pTop *Node = nil
	for {
		if pCurr == nil {
			break
		}
		pTemp := pCurr.Next
		pCurr.Next = pTop
		pTop = pCurr
		pCurr = pTemp
	}
	// l.PrintList()

}
