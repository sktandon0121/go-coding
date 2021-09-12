package main

func productOfArray(arr []int) int {
	if len(arr) == 0 {
		return 1
	}

	return arr[0] * productOfArray(arr[1:])
}
