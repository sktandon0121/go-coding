package main

import "fmt"

func main() {
	a := []int{2, 50}
	b := []int{3, 6, 18, 21}
	sortedArray := mergeSortedArray(a, b)

	fmt.Println(sortedArray)
}

func mergeSortedArray(A1 []int, A2 []int) []int {
	S := []int{}

	if len(A1) == 0 {
		S = A2
	} else if len(A2) == 0 {
		S = A1
	}

	item1, item2, i, j := A1[0], A2[0], 1, 1

	for i < len(A1) || j < len(A2) {
		fmt.Println(i, j)
		if item1 < item2 {
			S = append(S, item1)
			item1 = A1[i]
			i++
		} else {
			S = append(S, item2)
			item2 = A2[j]
			j++
		}
	}

	return S
}
