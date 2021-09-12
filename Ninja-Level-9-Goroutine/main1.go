package main

import "fmt"

func findA(num int) int {
	a := 100 //hidden
	if num > a {
		return 1
	} else if num < a {
		return -1
	} else {
		return 0
	}
}

func main() {
	a := 0
	if findA(a) == a {
		fmt.Println(a)
		return
	}
	a = 1
	i := 1
	for a != 0 {
		a = findA(i)
		if a > 0 {
			i--
		}
		if a < 0 {
			i++
		}
	}
	fmt.Println(i)
}
