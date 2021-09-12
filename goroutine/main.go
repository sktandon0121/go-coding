package main

import "fmt"

func main() {
	fmt.Println("main method ")

	c := make(chan int)

	for i := 1; i <= 10; i++ {
		go printNum(i, c)
		fmt.Println(<-c)
		i++
	}

}

func printNum(i int, c chan int) {
	fmt.Println(i)
	i++
	c <- i
}
