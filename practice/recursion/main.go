package main

import (
	"encoding/json"
	"fmt"
)

var r []byte

func main() {
	fmt.Println("main method start..........")

	// str := []string{"subodh", "manu", "harshu", "sktan", "hari"}

	head := &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}}
	// a := makeList(head)

	// head1 := &ListNode{}
	newList := removeElements(head, 7)

	b, _ := json.MarshalIndent(newList, " ", " ")

	fmt.Println(string(b))
	// fmt.Println(a)

}
