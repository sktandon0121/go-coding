package main

func sumOfDigit(num int) int {
	if num == 0 {
		return num
	}

	rem := num % 10
	q := num / 10

	return rem + sumOfDigit(q)

}
