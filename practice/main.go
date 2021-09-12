package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := true
	b := 5
	fmt.Println(swap(a, b))
	fmt.Println(a, b)
}

func swap(a interface{}, b interface{}) (interface{}, interface{}) {
	fmt.Println(reflect.TypeOf(a))
	return b, a
}
