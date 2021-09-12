package list

import "fmt"

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		val: 2,
		next: &MyLinkedList{val: 3,
			next: &MyLinkedList{val: 4, next: nil},
		},
	}
}

func (l *MyLinkedList) PrintList() {
	head := l
	count := 1
	for head != nil {

		fmt.Printf("%d ", head.val)
		count++
		head = head.next
	}
}

func (l *MyLinkedList) Get(index int) int {
	res := -1
	if index < 0 {
		return res
	}

	node := l
	i := 0
	for node.next != nil {
		if i == index {
			res = node.val
			break
		}
		i++
		node = node.next
	}

	return res
}
