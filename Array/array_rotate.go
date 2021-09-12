package main

import "fmt"

func main() {

	arr := []int{1, 2, 4, 3, 5, 7, 9}

	rotatedArr := array_rotate(arr, 3)

	fmt.Println(rotatedArr)
}

func array_rotate(arr []int, pos int) []int {
	a := make([]int, 3)
	b := make([]int, 3)

	if len(arr) == 0 {
		return arr
	} else {
		a = arr[0:pos]
		b = arr[pos:]

		for _, it := range a {
			b = append(b, it)
		}
	}

	return b
}
