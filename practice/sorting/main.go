package main

import "fmt"

func main() {
	fmt.Println("main method start..........")

	arr := []int{5, 4, 7, 1, 6}
	fmt.Println(arr)

	a := bubbleSort(arr)

	fmt.Println(a)
}
