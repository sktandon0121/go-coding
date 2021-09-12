package main

import "fmt"

func main1() {
	num := 123

	rem := num
	count := 0
	for rem != 0 {
		count++
		num = num / 10
		rem = num % 10
	}

	fmt.Println(count)

	st := &Stack{}

	st.Push("subodh")
	st.Push("Hello")

	fmt.Printf("Value in stack %#v \n", st)

	fmt.Println(st.Pop(), st.Pop())

}

type Stack struct {
	Value []string
}

func (st *Stack) Push(value string) *Stack {
	if value == "" {
		return nil
	}
	st.Value = append(st.Value, value)

	return st
}

func (st *Stack) Pop() string {
	if len(st.Value) == 0 {
		return ""
	}

	str := st.Value[len(st.Value)-1:]
	st.Value = st.Value[:len(st.Value)-1]
	return str[0]

}

func (st *Stack) PrintStack() {
	if len(st.Value) != 0 {
		for _, v := range st.Value {
			fmt.Printf("%s ", v)
		}
	}
}
