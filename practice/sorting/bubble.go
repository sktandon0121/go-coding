package main

func bubbleSort(arr []int) []int {

	for j := 0; j <= len(arr)-1; j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				swap(&arr[i], &arr[i+1])
			}
		}
	}

	return arr
}

func swap(a *int, b *int) {
	// temp := *b
	// *b = *a
	// *a = temp
	*a, *b = *b, *a
}
