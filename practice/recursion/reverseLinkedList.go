package main

/* *
 * Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(head *ListNode) *ListNode {
	// [1,2,6,3,4,5,6]
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 6}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 5}
	node7 := &ListNode{Val: 6}

	head = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	return head
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} /*  else if head.Next == nil {
		if head.Val == val {
			return nil
		} else {
			return head
		}
	} */

	preNode := head
	nextNode := head
	for nextNode != nil {
		if nextNode.Val == val {
			preNode.Next = nextNode.Next
		}
		preNode = nextNode
		nextNode = nextNode.Next
	}
	return head
}

func removeFromFirst(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	head = head.Next

	return head
}

func removeFromLast(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	preNode := head
	nextNode := head.Next

	for nextNode.Next != nil {
		preNode = nextNode
		nextNode = nextNode.Next
	}
	preNode.Next = nil
	return head
}
