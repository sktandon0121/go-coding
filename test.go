package main

import "fmt"

func test() {
	fmt.Println("Hello")
	go SayHello()
}

func SayHello() {
	fmt.Println("Hello Subodh")
}
